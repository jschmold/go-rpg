package game

type Entity struct {
	id     string
	health int32
}

type MagicEntity struct {
	Entity
	mana int32
}

func NewEntity(id string, health int32) Entity {
	return Entity{id, health}
}

func NewMagicEntity(id string, health int32, mana int32) MagicEntity {
	ent := Entity{id, health}
	return MagicEntity{ent, mana}
}

func (e *Entity) Damage(amt int32) int32 {
	e.health -= amt
	return e.GetHealth()
}

func (e *Entity) Heal(amt int32) int32 {
	e.health += amt
	return e.GetHealth()
}

func (e *Entity) GetHealth() int32 {
	return e.health
}

func (e *MagicEntity) GetMana() int32 {
	return e.mana
}

func (e *MagicEntity) RestoreMana(amt int32) int32 {
	e.mana += amt
	return e.GetMana()
}

func (e *MagicEntity) UseMana(amt int32) int32 {
	e.mana -= amt
	return e.GetMana()
}
