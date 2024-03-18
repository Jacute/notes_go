package main

type Note struct {
	name        string // Название записи
	description string // Описание записи
	author      string // Уникальный идентификатор автора записи
	is_private  bool   // Флаг, указывающий, является ли запись приватной или нет
}
