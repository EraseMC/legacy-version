package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// EnchantmentOption represents a single option in the enchantment table for a single item.
type EnchantmentOption struct {
	// Cost is the cost of the option. This is the amount of XP levels required to select this enchantment
	// option.
	Cost uint8
	// Enchantments holds the enchantments that will be applied to the item when this option is clicked.
	Enchantments protocol.ItemEnchantments
	// Name is a name that will be translated to the 'Standard Galactic Alphabet'
	// (https://minecraft.wiki/w/Enchanting_Table#Standard_Galactic_Alphabet) client-side. The names
	// generally have no meaning, such as:
	// 'animal imbue range galvanize '
	// 'bless inside creature shrink '
	// 'elder free of inside '
	Name string
	// RecipeNetworkID is a unique network ID for this enchantment option. When enchanting, the client
	// will submit this network ID in a ItemStackRequest packet with the CraftRecipe action, so that the
	// server knows which enchantment was selected.
	// Note that this ID should still be unique with other actual recipes. It's recommended to start counting
	// for enchantment network IDs from the counter used for producing network IDs for the normal recipes.
	RecipeNetworkID uint32
}

// Marshal encodes/decodes an EnchantmentOption.
func (x *EnchantmentOption) Marshal(r protocol.IO) {
	r.Uint8(&x.Cost)
	protocol.Single(r, &x.Enchantments)
	r.String(&x.Name)
	r.Varuint32(&x.RecipeNetworkID)
}

// ToLatest ...
func (x *EnchantmentOption) ToLatest() protocol.EnchantmentOption {
	return protocol.EnchantmentOption{
		Cost:            x.Cost,
		Enchantments:    x.Enchantments,
		Name:            x.Name,
		RecipeNetworkID: x.RecipeNetworkID,
	}
}

// FromLatest ...
func (x *EnchantmentOption) FromLatest(latest protocol.EnchantmentOption) EnchantmentOption {
	x.Cost = latest.Cost
	x.Enchantments = latest.Enchantments
	x.Name = latest.Name
	x.RecipeNetworkID = latest.RecipeNetworkID
	return *x
}
