package model

func TransformEmptyRoom(raw *EmptyRoomRawInfo) EmptyRoomInfo {
	var emptyRoomInfo EmptyRoomInfo

	for _, value := range raw.Items {
		emptyRoomInfo = append(emptyRoomInfo,
			&Room{
				RoomName:         value.Cdmc,
				BuildName:        value.Jxlmc,
				RoomSize:         value.Jzmj,
				RoomSeats:        value.Zws,
				RoomSeatsForExam: value.Kszws1,
				Campus:           value.Xqmc,
				Type:             value.Cdlbmc,
			})
	}
	return emptyRoomInfo
}

type Room struct {
	RoomName         string `json:"roomName"`
	BuildName        string `json:"buildName"`
	RoomSize         string `json:"roomSize"`
	RoomSeats        string `json:"roomSeats"`
	RoomSeatsForExam string `json:"roomSeatsForExam"`
	Campus           string `json:"campus"`
	Type             string `json:"type"`
}

type EmptyRoomInfo []*Room

type EmptyRoomRawInfo struct {
	Items []*struct {
		Bz     string
		Cdlbmc string
		Cdmc   string
		Jxlmc  string
		Jzmj   string
		Kszws1 string
		Xqmc   string
		Zws    string
	}
}
