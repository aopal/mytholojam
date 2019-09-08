testMove = require("moves/test")
defaults = require("defaults")

local _M = {}

function _M.load(filename, ctx)
  t = table.shallow_copy(testMove)
  t._type="move_template"
  table.insert(t.path_chosen_hooks, defaults.OnPathChosen)
  table.insert(t.move_finished_hooks, defaults.OnMoveFinished)
  table.insert(t.equipment_hit_hooks, defaults.OnEquipmentHit)
  return t
end


return _M