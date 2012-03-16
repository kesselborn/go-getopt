if &cp | set nocp | endif
let s:cpo_save=&cpo
set cpo&vim
inoremap <silent> <Plug>allmlXmlV ="&#".getchar().";"
map! <D-v> *
map  "+y
map 	 :Bi
map  :wincmd W
map Q gq
xmap S <Plug>VSurround
vmap [% [%m'gv``
nnoremap <silent> \b :CommandTBuffer
nnoremap <silent> \t :CommandT
nmap \cv <Plug>VCSVimDiff
nmap \cu <Plug>VCSUpdate
nmap \cU <Plug>VCSUnlock
nmap \cr <Plug>VCSReview
nmap \cq <Plug>VCSRevert
nmap \cl <Plug>VCSLog
nmap \cL <Plug>VCSLock
nmap \ci <Plug>VCSInfo
nmap \cg <Plug>VCSGotoOriginal
nmap \cd <Plug>VCSDiff
nmap \cD <Plug>VCSDelete
nmap \cc <Plug>VCSCommit
nmap \cG <Plug>VCSClearAndGotoOriginal
nmap \cn <Plug>VCSAnnotate
nmap \ca <Plug>VCSAdd
map \T <Plug>RubyFileRun
map \r <Plug>RubyTestRun
nnoremap <silent> \ff :call g:Jsbeautify()
nmap \so <Plug>DBOrientationToggle
nmap \sh <Plug>DBHistory
nmap \slv <Plug>DBListView
nmap \slp <Plug>DBListProcedure
nmap \slt <Plug>DBListTable
vmap <silent> \slc :exec 'DBListColumn '.DB_getVisualBlock()
nmap \slc <Plug>DBListColumn
nmap \sbp <Plug>DBPromptForBufferParameters
nmap \sdpa <Plug>DBDescribeProcedureAskName
vmap <silent> \sdp :exec 'DBDescribeProcedure '.DB_getVisualBlock()
nmap \sdp <Plug>DBDescribeProcedure
nmap \sdta <Plug>DBDescribeTableAskName
vmap <silent> \sdt :exec 'DBDescribeTable '.DB_getVisualBlock()
nmap \sdt <Plug>DBDescribeTable
vmap <silent> \sT :exec 'DBSelectFromTableTopX '.DB_getVisualBlock()
nmap \sT <Plug>DBSelectFromTableTopX
nmap \sta <Plug>DBSelectFromTableAskName
nmap \stw <Plug>DBSelectFromTableWithWhere
vmap <silent> \st :exec 'DBSelectFromTable '.DB_getVisualBlock()
nmap \st <Plug>DBSelectFromTable
nmap <silent> \sel :.,.DBExecRangeSQL
nmap <silent> \sea :1,$DBExecRangeSQL
nmap \sE <Plug>DBExecSQLUnderCursorTopX
nmap \se <Plug>DBExecSQLUnderCursor
vmap \sE <Plug>DBExecVisualSQLTopX
vmap \se <Plug>DBExecVisualSQL
map \rwp <Plug>RestoreWinPosn
map \swp <Plug>SaveWinPosn
nmap \s :call InitShell()
map \cs :!git status
vmap ]% ]%m'gv``
vmap a% [%v]%
nmap cs <Plug>Csurround
nmap ds <Plug>Dsurround
nmap gx <Plug>NetrwBrowseX
xmap s <Plug>Vsurround
nmap ySS <Plug>YSsurround
nmap ySs <Plug>YSsurround
nmap yss <Plug>Yssurround
nmap yS <Plug>YSurround
nmap ys <Plug>Ysurround
nnoremap <silent> <Plug>NetrwBrowseX :call netrw#NetrwBrowseX(expand("<cWORD>"),0)
nnoremap <silent> <Plug>CVSWatchRemove :CVSWatch remove
nnoremap <silent> <Plug>CVSWatchOn :CVSWatch on
nnoremap <silent> <Plug>CVSWatchOff :CVSWatch off
nnoremap <silent> <Plug>CVSWatchAdd :CVSWatch add
nnoremap <silent> <Plug>CVSWatchers :CVSWatchers
nnoremap <silent> <Plug>CVSUnedit :CVSUnedit
nnoremap <silent> <Plug>CVSEditors :CVSEditors
nnoremap <silent> <Plug>CVSEdit :CVSEdit
nnoremap <silent> <Plug>VCSVimDiff :VCSVimDiff
nnoremap <silent> <Plug>VCSUpdate :VCSUpdate
nnoremap <silent> <Plug>VCSUnlock :VCSUnlock
nnoremap <silent> <Plug>VCSStatus :VCSStatus
nnoremap <silent> <Plug>VCSReview :VCSReview
nnoremap <silent> <Plug>VCSRevert :VCSRevert
nnoremap <silent> <Plug>VCSLog :VCSLog
nnoremap <silent> <Plug>VCSLock :VCSLock
nnoremap <silent> <Plug>VCSInfo :VCSInfo
nnoremap <silent> <Plug>VCSClearAndGotoOriginal :VCSGotoOriginal!
nnoremap <silent> <Plug>VCSGotoOriginal :VCSGotoOriginal
nnoremap <silent> <Plug>VCSDiff :VCSDiff
nnoremap <silent> <Plug>VCSDelete :VCSDelete
nnoremap <silent> <Plug>VCSCommit :VCSCommit
nnoremap <silent> <Plug>VCSAnnotate :VCSAnnotate
nnoremap <silent> <Plug>VCSAdd :VCSAdd
nmap <silent> <Plug>RestoreWinPosn :call RestoreWinPosn()
nmap <silent> <Plug>SaveWinPosn :call SaveWinPosn()
map <F2> :let g:rubytest_in_quickfix=
vmap <BS> "-d
vmap <D-x> "*d
vmap <D-c> "*y
vmap <D-v> "-d"*P
nmap <D-v> "*P
imap S <Plug>ISurround
imap s <Plug>Isurround
imap  <Plug>Isurround
let &cpo=s:cpo_save
unlet s:cpo_save
set autowrite
set backspace=indent,eol,start
set backup
set backupdir=/tmp
set expandtab
set fileencodings=utf-8
set helplang=en
set history=50
set hlsearch
set incsearch
set laststatus=2
set ruler
set runtimepath=~/.vim,~/.vim/bundle/command-t,~/.vim/bundle/nerdtree,~/.vim/bundle/vim-fugitive,~/.vim/bundle/vim-powerline,/usr/local/Cellar/vim/7.3.333/share/vim/vimfiles,/usr/local/Cellar/vim/7.3.333/share/vim/vim73,/usr/local/Cellar/vim/7.3.333/share/vim/vimfiles/after,~/.vim/after
set scrolloff=12
set shiftwidth=2
set showcmd
set softtabstop=2
set noswapfile
set tabstop=2
set tags=./tags,tags,$GEM_HOME/tags
set updatetime=1000
set wildmenu
set wildmode=longest,list,full
" vim: set ft=vim :
