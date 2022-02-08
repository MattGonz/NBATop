# NBATop
A [gocui](https://github.com/jroimartin/gocui) terminal UI for NBA stats (with [vim](https://en.wikipedia.org/wiki/Vim_(text_editor))-like keybinds)

#### Demo gif
<img src="demo.gif" width="833" height="492"/>

### Build from source
```
go get github.com/MattGonz/NBATop
cd $GOPATH/src/github.com/MattGonz/NBATop
go build -o nbatop main.go
./nbatop
```




#### Todo

###### General
- [ ] Config file
- [ ] Cache requests

###### Views
- [x] Games today
- [x] Games today -> box scores
- [x] Standings
- [x] Team game logs
- [x] Game box scores
- [x] Player game logs
- [x] Horizontal scrolling
- [ ] Player profiles
- [ ] Multiple seasons
- [ ] `?` → help view


###### UI improvements
- [ ] UI formatting keybinds
- [ ] Table sorting
- [ ] Fuzzing finding / better navigation keybinds





##### Inspirations
* [jesseduffield/lazygit](https://github.com/jesseduffield/lazygit) and [lazydocker](https://github.com/jesseduffield/lazydocker)
* [ClementTsang/bottom](https://github.com/ClementTsang/bottom)
* [miguelmota/cointop](https://github.com/cointop-sh/cointop)
* The [Unofficial NBA API documentation](http://nbasense.com/nba-api/) of [jasonroman/nba-api](https://github.com/jasonroman/nba-api)
