local defaults = require("defaults")
local uuid = require("uuid")
local MT = require("move_template")

local str = uuid.new()

function love.draw()
  local mt = MT.load("asdf",  {})
  love.graphics.print(mt.move_finished_hooks[1](), 400, 300)
end




function unravel_callbacks(ctx, cb_arr)
  for i, func in ipairs(cb_arr) do 
    func(ctx)
  end
end

function calculate_damage(ctx)
  return ctx.attacked_with.power / (ctx.defender.base_defense * ctx.defender.defense_mult)
end


function table.shallow_copy(t)
  local t2 = {}
  for k,v in pairs(t) do
    t2[k] = v
  end
  return t2
end