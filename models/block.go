package models

import "autumn/pkg/autumncore/generic"

type Block struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	GroupLocationID int    `json:"id_group_location"`
	*generic.ModelImpl[Block]
}

func (b *Block) TableName() string {
	return "m_blok"
}
