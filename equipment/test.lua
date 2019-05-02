testMove=require("moves/test")

return {
  name="testEquip",
  type="slash",
  moves=[testMove]
  max_health=100
  base_defense=10

  hit_hooks={},
  damage_taken_hooks={},
  death_hooks={}
}