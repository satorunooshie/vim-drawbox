" -------- Keyboard mappings --------

" -------- Box drawing --------
" rectangle
vnoremap +o :<C-u>call drawbox#draw("+o", [])<CR>
" labeled rectangle
vnoremap +O :<C-u>call drawbox#draw_with_label("+O", [])<CR>
vnoremap +mcb :<C-u>call drawbox#draw_with_label("+mcb", [])<CR>

" middle left labeled rectangle
vnoremap +[O :<C-u>call drawbox#draw_with_label("+[O", [])<CR>
vnoremap mlb :<C-u>call drawbox#draw_with_label("+mlb", [])<CR>
" middle right labeled rectangle
vnoremap +]O :<C-u>call drawbox#draw_with_label("+]O", [])<CR>
vnoremap +mrb :<C-u>call drawbox#draw_with_label("+mrb", [])<CR>
" top center labeled rectangle
vnoremap +{[O :<C-u>call drawbox#draw_with_label("+{[O", [])<CR>
vnoremap +tcb :<C-u>call drawbox#draw_with_label("+tcb", [])<CR>
" top right labeled rectangle
vnoremap +{]O :<C-u>call drawbox#draw_with_label("+{]O", [])<CR>
vnoremap +trb :<C-u>call drawbox#draw_with_label("+trb", [])<CR>
" bottom center labeled rectangle
vnoremap +}[O :<C-u>call drawbox#draw_with_label("+}[O", [])<CR>
vnoremap +bcb :<C-u>call drawbox#draw_with_label("+bcb", [])<CR>
" bottom right labeled rectangle
vnoremap +}]O :<C-u>call drawbox#draw_with_label("+}]O", [])<CR>
vnoremap +brb :<C-u>call drawbox#draw_with_label("+brb", [])<CR>

" -------- Labeling --------
" middle center
vnoremap +c :<C-u>call drawbox#draw_with_label("+c", [])<CR>
vnoremap +mcl :<C-u>call drawbox#draw_with_label("+mcl", [])<CR>
" middle left
vnoremap +[c :<C-u>call drawbox#draw_with_label("+[c", [])<CR>
vnoremap +mll :<C-u>call drawbox#draw_with_label("+mll", [])<CR>
" middle right
vnoremap +]c :<C-u>call drawbox#draw_with_label("+]c", [])<CR>
vnoremap +mrl :<C-u>call drawbox#draw_with_label("+mrl", [])<CR>

" top center
vnoremap +{c :<C-u>call drawbox#draw_with_label("+{c", [])<CR>
vnoremap +tcl :<C-u>call drawbox#draw_with_label("+tcl", [])<CR>
" top left
vnoremap +{[c :<C-u>call drawbox#draw_with_label("+{[c", [])<CR>
vnoremap +tll :<C-u>call drawbox#draw_with_label("+tll", [])<CR>
" top right
vnoremap +{]c :<C-u>call drawbox#draw_with_label("+{]c", [])<CR>
vnoremap +trl :<C-u>call drawbox#draw_with_label("+trl", [])<CR>

" bottom center
vnoremap +}c :<C-u>call drawbox#draw_with_label("+}c", [])<CR>
vnoremap +bcl :<C-u>call drawbox#draw_with_label("+bcl", [])<CR>
" bottom left
vnoremap +}[c :<C-u>call drawbox#draw_with_label("+}[c", [])<CR>
vnoremap +bll :<C-u>call drawbox#draw_with_label("+bll", [])<CR>
" bottom right
vnoremap +}]c :<C-u>call drawbox#draw_with_label("+}]c", [])<CR>
vnoremap +brl :<C-u>call drawbox#draw_with_label("+brl", [])<CR>

" -------- Line drawing --------
vnoremap +> :<C-u>call drawbox#draw("+>", [])<CR>
vnoremap +< :<C-u>call drawbox#draw("+<", [])<CR>

vnoremap +v :<C-u>call drawbox#draw("+v", [])<CR>
vnoremap +V :<C-u>call drawbox#draw("+v", [])<CR>
vnoremap +^ :<C-u>call drawbox#draw("+^", [])<CR>

vnoremap ++> :<C-u>call drawbox#draw("++>", [])<CR>
vnoremap ++< :<C-u>call drawbox#draw("++<", [])<CR>

vnoremap ++v :<C-u>call drawbox#draw("++v", [])<CR>
vnoremap ++V :<C-u>call drawbox#draw("++v", [])<CR>
vnoremap ++^ :<C-u>call drawbox#draw("++^", [])<CR>

vnoremap +- :<C-u>call drawbox#draw("+-", [])<CR>
vnoremap +_ :<C-u>call drawbox#draw("+_", [])<CR>

vnoremap +\| :<C-u>call drawbox#draw("+\|", [])<CR>

" -------- Selection --------
vnoremap ao :<C-u>call drawbox#select("ao")<CR>
vnoremap ab :<C-u>call drawbox#select("ao")<CR>
vnoremap io :<C-u>call drawbox#select("io")<CR>
vnoremap ib :<C-u>call drawbox#select("io")<CR>
