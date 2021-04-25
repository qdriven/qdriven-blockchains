package io.qdriven.chains.config.tokens.entity;

import lombok.Data;
import xyz.erupt.annotation.Erupt;

import javax.persistence.Entity;
import javax.persistence.Table;
//type Chain struct {
//        ChainId             *uint64 `gorm:"primaryKey;type:bigint(20);not null"`
//        Height              uint64  `gorm:"type:bigint(20);not null"`
//        HeightSwap          uint64  `gorm:"type:bigint(20);not null"`
//        BackwardBlockNumber uint64  `gorm:"type:bigint(20);not null"`
//        }
/**
 * @author Patrick
 **/
@Erupt(name = "chains")
@Entity
@Table(name = "chains")
@Data
public class Blockchain {

}
