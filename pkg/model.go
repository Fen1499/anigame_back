package pkg

type base struct {
	Id string
    InsertedAt string
    UpdatedAt string
}

type daily struct {
    base
    Romanji string
    ImgURL string
    AnilistId string
    PickedOn string
}

type UserDaily struct {
    base
    DailyId string
    Tries string //lista de letras
    Done bool
}
