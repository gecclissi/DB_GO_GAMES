package models

type Calcula struct {
	IDAuxCalcula        int `json:"id_aux_calcula" gorm:"primaryKey"`
	IdRespostaRealizada int `json:"id_resposta_realizada"`
	IdContaPonto        int `json:"id_conta_ponto"`
}

type CalculaInput struct {
	IdRespostaRealizada int `json:"id_resposta_realizada" binding:"required"`
	IdContaPonto        int `json:"id_conta_ponto" binding:"required"`
}
