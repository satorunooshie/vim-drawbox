let s:cmd = 'drawbox'
if !executable(s:cmd)
    echohl ErrorMsg
    echomsg 'not found executable ' . s:cmd
    echohl None
finish
endif

function! drawbox#get_end_pos()
	" Vim reports '< and '> in the wrong order if the end of the selection
	" is in an earlier line than the start of the selection. This is why
	" we need this hack.
	let m = getpos("'m")
	execute "normal! gvmm\<Esc>"
	let p = getpos("'m")
	call setpos("'m", m)
	return p
endfunction

function! drawbox#get_start_pos(startPos)
	" Returns the 'other corner' of the visual selection.
	let p1 = getpos("'<")
	let p2 = getpos("'>")
	if p1 == a:startPos
		return p2
	endif
    return p1
endfunction

function! drawbox#draw(cmd, args)
	let p2 = drawbox#get_end_pos()
	let p1 = drawbox#get_start_pos(p2)
	let y1 = p1[1] - 1
	let y2 = p2[1] - 1
	let x1 = p1[2] + p1[3] - 1
	let x2 = p2[2] + p2[3] - 1
	let c = [s:cmd, shellescape(a:cmd), y1, x1, y2, x2] + a:args
	execute "%!" . join(c, " ")
	call setpos(".", p2)
endfunction

function! drawbox#draw_with_label(cmd, args)
	let label = shellescape(input("Label: "))
	call drawbox#draw(a:cmd, [label] + a:args)
endfunction

function! drawbox#select(cmd)
	let p2 = drawbox#get_end_pos()
	let p1 = drawbox#get_start_pos(p2)
	let y1 = p1[1] - 1
	let y2 = p2[1] - 1
	let x1 = p1[2] + p1[3] - 1
	let x2 = p2[2] + p2[3] - 1

	let contents = join(getline(1,'$'), "\n")
	let c = [s:cmd, shellescape(a:cmd), y1, x1, y2, x2]
	let result = system(join(c, " "), contents)

	let coords = split(result, ",")

	call setpos("'<", [0, coords[0]+1, coords[1]+1, 0])
	call setpos("'>", [0, coords[2]+1, coords[3]+1, 0])
	normal! gv
endfunction
