package game

import "testing"

func TestEntity(test *testing.T) {
	test.Run("Heal", func(t *testing.T) {
		ent := NewEntity("player", 100)
		health := ent.Heal(50)

		if health != 150 {
			t.Errorf("expect heath 150, actually %d", health)
		}

		if ent.health != 150 {
			t.Errorf("expect ent.heath 150, actually %d", health)
		}
	})

	test.Run("Damage", func(t *testing.T) {
		ent := NewEntity("player", 100)
		health := ent.Damage(50)

		if ent.health != 50 {
			t.Errorf("expected ent.health 50, actually %d", health)
		}

		if health != 50 {
			t.Errorf("expected health 50, actually %d", health)
		}
	})
}

func TestMagicEntity(test *testing.T) {
	test.Run("RestoreMana", func(t *testing.T) {
		ent := NewMagicEntity("player", 100, 30)
		mana := ent.RestoreMana(20)

		if mana != 50 {
			t.Errorf("expected mana 50, actually %d", mana)
		}

		if ent.mana != 50 {
			t.Errorf("expected ent.mana 50, actually %d", ent.mana)
		}
	})

	test.Run("UseMana", func(t *testing.T) {
		ent := NewMagicEntity("player", 100, 30)

		mana := ent.UseMana(30)

		if mana != 0 {
			t.Errorf("expected mana to be 0, actually %d", mana)
		}

		if ent.mana != 0 {
			t.Errorf("expected ent.mana to be 0, actually %d", ent.mana)
		}

	})
}
