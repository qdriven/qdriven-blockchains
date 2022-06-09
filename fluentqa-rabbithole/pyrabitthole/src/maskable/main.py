#!/usr/bin/env python
# -*- coding:utf-8 -*-
# This is a sample Python script.

# Press Shift+F10 to execute it or replace it with your code.
# Press Double Shift to search everywhere for classes, files, tool windows, actions, and settings.
import web
import sqlite3
import json
from urllib.request import Request, urlopen


def customhook():
    web.header('Access-Control-Allow-Origin', '*')
    web.header('Access-Control-Allow-Headers',
               'Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers, X-Requested-By, Access-Control-Allow-Methods')
    web.header('Access-Control-Allow-Methods', 'POST, GET, PUT, DELETE')


class index:
    def GET(self):
        conn = sqlite3.connect("./test.db")
        cur = conn.cursor()
        return "hello"


class badge:
    def GET(self):
        # print(params)
        json_data = web.input()
        wallet_address = json_data['wallet_address']
        twitter_user_id = json_data['twitter_user_id']
        twitter_avatar = json_data['twitter_avatar']
        twitter_nickname = json_data['twitter_nickname']
        # 处理twitter
        add_twitter(twitter_user_id, twitter_avatar, twitter_nickname)
        # 处理绑定关系
        add_bind(twitter_user_id, wallet_address)
        # 处理nft信息
        res = add_nft(wallet_address)
        response = {
            'code': 0,
            'message': 'ok',
            'data': res
        }
        return response


class twitter:
    def GET(self):
        params = web.input(badge_address=None)
        badge_address = params.badge_address
        conn = sqlite3.connect("./test.db")
        cur = conn.cursor()
        sql1 = "select badge_code,badge_name,badge_icon,badge_address from badge where badge_address ='{0}' and is_delete=0".format(
            badge_address)
        cur.execute(sql1)
        badge_res = cur.fetchone()
        sql2 = '''
           SELECT  ti.twitter_user_id,twitter_avatar,twitter_nickname FROM user_badge u
           INNER  JOIN twitter_bind t on t.wallet_address= u.wallet_address
           INNER JOIN  twitter_info ti on ti.twitter_user_id = t.twitter_user_id
           WHERE u.contract_address ='{0}'
        '''.format(badge_address)
        cur.execute(sql2)
        results = cur.fetchall()
        res = {}
        if badge_res is not None:
            res['badge_address'] = badge_address
            res['badge_name'] = badge_res[1]
            res['badge_icon'] = badge_res[2]
            res['badge_code'] = badge_res[0]
            res['tts'] = []
        if results is not None:
            users = ({'twitter_user_id': i[0],
                      'twitter_avatar': i[1],
                      'twitter_nickname': i[2]
                      } for i in results)
            res['tts'] = list(users)
        res_result = {
            'code': 0,
            'message': 'ok',
            'data': res
        }
        return res_result


def add_twitter(twitter_user_id, twitter_avatar, twitter_nickname):
    conn = sqlite3.connect("./test.db")
    sql1 = "select count(1) from twitter_info where twitter_user_id ='{0}'".format(twitter_user_id)
    cur = conn.cursor()
    cur.execute(sql1)
    results = cur.fetchone()
    count = results[0]
    if count == 0:
        sql2 = "insert into twitter_info(twitter_user_id,twitter_avatar,twitter_nickname)values('{0}','{1}','{2}')".format(
            twitter_user_id, twitter_avatar, twitter_nickname)
        cur.execute(sql2)
    else:
        sql3 = "update twitter_info set twitter_avatar='{0}',twitter_nickname='{1}' where twitter_user_id='{2}'".format(
            twitter_avatar, twitter_nickname, twitter_user_id)
        cur.execute(sql3)
    conn.commit()
    cur.close()
    conn.close()


def add_bind(twitter_user_id, wallet_address):
    conn = sqlite3.connect("./test.db")
    sql1 = "select count(1) from twitter_bind where is_delete=0 and twitter_user_id='{0}'".format(twitter_user_id)
    cur = conn.cursor()
    cur.execute(sql1)
    results = cur.fetchone()
    count = results[0]
    if count == 0:
        sql2 = "insert into twitter_bind(twitter_user_id,wallet_address)values('{0}','{1}')".format(
            twitter_user_id, wallet_address)
        cur.execute(sql2)
    else:
        sql3 = "update twitter_bind set wallet_address='{0}' where twitter_user_id='{1}'".format(wallet_address,
                                                                                                 twitter_user_id)
        cur.execute(sql3)
    conn.commit()
    cur.close()
    conn.close()


def add_nft(wallet_address):
    sql1 = "select badge_address,badge_name,badge_icon,badge_code from badge where is_delete=0"
    conn = sqlite3.connect("./test.db")
    cur = conn.cursor()
    cur.execute(sql1)
    results = cur.fetchall()
    badges = {}
    sub_url = ""
    for r in results:
        badge_code = r[3]
        address = r[0]
        badges[address] = r[1] + "$@" + r[2] + "$@" + badge_code
        sub_url = sub_url + "&asset_contract_addresses=" + address
    url = "https://api.opensea.io/api/v1/assets?owner={0}{1}".format(wallet_address, sub_url)
    print(url)
    req = Request(url, headers={'User-Agent': 'Mozilla/5.0'})
    resp = urlopen(req)
    web_byte = resp.read()
    webpage = web_byte.decode('utf-8')
    jsons = json.loads(webpage)
    result = []
    contracts = {}
    for asset in jsons['assets']:
        contract = asset['asset_contract']['address']
        name = asset['name']
        badge = badges[contract]
        badge_sub = badge.split("$@")
        if contract not in contracts:
            contracts[contract] = 1
            result.append({
                "badge_address": contract,
                "badge_name": badge_sub[0],
                "badge_icon": badge_sub[1],
                "badge_code": badge_sub[2]
            })
    add_one_badge(wallet_address, contracts)
    return result


def add_one_badge(wallet_address, contracts):
    # 把不属于用户的徽章删除
    placeholders = ','.join("'{0}'".format(i) for i in contracts.keys())
    sql1 = "update user_badge set is_delete = 1 where wallet_address='{0}' and contract_address not in ({1} )".format(
        wallet_address, placeholders)
    conn = sqlite3.connect("./test.db")
    cur = conn.cursor()
    cur.execute(sql1)
    # 重新添加
    for c in contracts.keys():
        sql2 = "select count(1) from user_badge where wallet_address='{0}' and contract_address ='{1}'".format(
            wallet_address, c)
        cur.execute(sql2)
        count = cur.fetchone()
        if count[0] == 0:
            sql3 = "insert into user_badge(wallet_address,contract_address) values('{0}','{1}')".format(wallet_address,
                                                                                                        c)
            cur.execute(sql3)
    conn.commit()
    cur.close()
    conn.close()


urls = ("/index", index,
        "/get_badges_by_address", badge,
        "/get_twitters_by_badge", twitter
        )

# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    app = web.application(urls, globals())
    app.add_processor(web.loadhook(customhook))
    app.run()
