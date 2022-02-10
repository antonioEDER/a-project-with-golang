package ordereds

// Filter ...
type Filter struct {
	ID                  int64  `query:"id"`
	Tipos_Pedidos_Ids   string `query:"tipos_pedidos_ids"`
	Tipos_Status_Ids    string `query:"tipos_status_ids"`
	Data_Inicio         string `query:"data_inicio"`
	Data_Final          string `query:"data_final"`
	Numero              int64  `query:"numero"`
	Visualizado         bool   `query:"visualizado"`
	Todos               bool   `query:"todos"`
	Parceiros_Id        string `query:"parceiros_id" json:"parceiros_id"`
	Pedidos_Id          int64  `query:"pedidos_id" json:"pedidos_id"`
	Pessoas_Id          int64  `query:"pessoas_id" json:"pessoas_id"`
	Pessoas_Usuarios_Id int64  `query:"pessoas_usuarios_id" json:"pessoas_usuarios_id"`
	Time_Zone           string `query:"time_zone" json:"time_zone"`
}
