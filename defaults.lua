local _M = {}

function _M.OnPathChosen(ctx)
  -- do nothing
end

function _M.OnMoveFinished(ctx)
  return "asdf"
  -- do nothing
end

function _M.OnEquipmentHit(ctx)
  -- do nothing
end

function _M.OnHit(ctx)
  local damage = calculate_damage(ctx)
  return _M.TakeDamage(ctx, damage)
end

function _M.OnDamageTaken(ctx)
  -- do nothing
    return "damaged"
end

function _M.OnDeath(ctx)
  return "dead"
end

function _M.TakeDamage(ctx, damage)
  ctx.defender.health -= damage

  if ctx.defender.health <= 0 then
    ctx.defender.health = 0
    return unravel_callbacks(ctx, ctx.defender.death_hooks)
  elseif ctx.defender.health > ctx.defender.max_health
    ctx.defender.health = ctx.defender.max_health
  end

  return unravel_callbacks(ctx.defender, ctx.defender.hit_hooks)
end

return _M