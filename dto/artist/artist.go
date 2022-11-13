package artistdto

type RequestCreateArtist struct {
	Name        string `json:"name"  validate:"required"`
	Old         int    `json:"old"  validate:"required"`
	Role        string `json:"role"  validate:"required"`
	StartCareer int    `json:"start_career"  validate:"required"`
}
