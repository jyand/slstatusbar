#!/usr/bin/env luajit

-- flags for lemonbar
--
bottom = false
-- empty strings for default values
lemonopts = {
        font = "\"-b&h-lucida-bold-r-normal-sans-12-120-75-75-p-79-iso10646-1\"",
        ul = "0", -- underline/overline size
        UC = "'#ffffff'", --underline color
        BG = "'#000000'", -- background color
        FG = "'#ffffff'" -- foreground color
}
geometry = {
        width = "",
        height = "24",
        xoffset = "0",
        yoffset = "0"
}

function BuildCmd(str, opts, bo)
        if bottom then str = str .. " -b" end
        -- str = str .. " -g "
        -- if #geom.width > 0 then str = str .. geom.width .. "x" end
        -- str = str .. geom.height .. "+" .. geom.xoffset .. "+" .. geom.yoffset
        for k, v in pairs(lemonopts) do
                if #v > 0 then
                        str = str .. " -" .. string.sub(k, 1, 1).. " " .. v
                end
        end
       str = str .. " &"
        return str
end

os.execute(BuildCmd("slstatus -s | lemonbar", lemonopts, bottom))

--print(BuildCmd("lemonbar", geometry, lemonopts, bottom))
