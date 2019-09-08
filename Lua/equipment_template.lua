testEquip = require("equipment/test")
defaults = require("defaults")

local _M = {}

function _M.load(filename, ctx)
  t = table.shallow_copy(testEquip)
  t._type="equipment_template"
  table.insert(t.hit_hooks, defaults.OnHit)
  table.insert(t.damage_taken_hooks, defaults.OnDamageTaken)
  table.insert(t.death_hooks, defaults.OnDeath)
  return t
end


return _M