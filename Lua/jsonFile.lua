-- load the JSON library.
local Json = require("json")

local JsonStorage = {}

-- Function to save a table.&nbsp; Since game settings need to be saved from session to session, we will
-- use the Documents Directory
JsonStorage.write = function(t, filename)
    -- local path = system.pathForFile( filename, system.DocumentsDirectory)
    local file = io.open(filename, "w")

    if file then
        local contents = Json.encode(t)
        file:write( contents )
        io.close( file )
        return true
    else
        return false
    end
end

JsonStorage.load = function(filename)
    -- local path = system.pathForFile( filename, system.DocumentsDirectory)
    local contents = ""
    local myTable = {}
    local file = io.open( filename, "r" )

    if file then
        -- read all contents of file into a string
        local contents = file:read( "*a" )
        myTable = Json.decode(contents);
        io.close( file )
        return myTable
    end
    return nil
end

return JsonStorage