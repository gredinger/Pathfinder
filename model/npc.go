package model

type NPC struct {
	Name    string
	HasWang bool
	spouse  []NPC
}

// GenerateRandomNPC generates a NPC randomly
func GenerateRandomNPC() NPC {
	return NPC{Name: "Gene", HasWang: true} // not very random
}

func AddSpose() {

}
