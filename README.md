我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

battery - Draw battery unicode art written by Go
=======
![sc_battery](https://cloud.githubusercontent.com/assets/6500104/19550024/6018c768-96e2-11e6-9ae1-f66b2406b8a7.png)  
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
  
battery unicode art on your tmux sessions or the terminal.  

## Status
(2019-11-03)  
Supported to display patched fonts.  
Very Thanks [TsutomuNakamura](https://github.com/TsutomuNakamura)!!

Let's run `battery -i` after installed patched fonts. [See more details](https://github.com/Code-Hex/battery#support-patched-fonts). 

(2018-05-01)  
Supported to show elapsed time.  
Very Thanks [delphinus](https://github.com/delphinus)!!

Let's run `battery -e`

![elapsed time](https://user-images.githubusercontent.com/1239245/39427036-388223ce-4cbd-11e8-859a-5363cdac3452.png)

(2016-10-24)  
linux supported.  
Thanks [mattn](https://github.com/mattn)!!  

(2016-10-21)  
windows supported.  
  
(2016-10-20)  
Now, this command can use mac user only.  
However, I hope to support ~~windows~~ and ~~linux~~, bsd in future.   
So, plz help me (´；ω；｀)  

## Installation
    go get -u github.com/Code-Hex/battery/cmd/battery

## Usage
For tmux user, please write `#(battery -t)` in your `.tmux.conf`  
Please refer to [this](https://github.com/Code-Hex/dotfiles/blob/master/tmux/.tmux.conf#L82)

### Support patched fonts
You can display the status of battery with patched fonts.
Installing the font, Inconsolata Nerd Font Complete.otf for example, you can do it like below.

* For mac
```
cd ~/Library/Fonts/
wget https://raw.githubusercontent.com/ryanoasis/nerd-fonts/master/patched-fonts/Inconsolata/complete/Inconsolata%20Nerd%20Font%20Complete.otf
```

* For Linux
```
cd ~/.local/share/fonts
wget https://raw.githubusercontent.com/ryanoasis/nerd-fonts/master/patched-fonts/Inconsolata/complete/Inconsolata%20Nerd%20Font%20Complete.otf
```

* For Windows
```
* Open your browser then fill this url
  https://raw.githubusercontent.com/ryanoasis/nerd-fonts/master/patched-fonts/Inconsolata/complete/Inconsolata%20Nerd%20Font%20Complete.otf
* Double click the file that you downloaded and the window will be opened
* Click the install button
```

Your environment might require to reboot the OS. If so, reboot your OS.
Then open your terminal and set your preferences to use it.

Then for tmux user, please write `#(battery -t -i)` in your `.tmux.conf`. 

![patched_font_100](https://user-images.githubusercontent.com/10674169/58262398-2d0ead80-7db5-11e9-816e-7df5a416aed2.png)
![patched_font_50](https://user-images.githubusercontent.com/10674169/58262403-2f710780-7db5-11e9-8a8c-e63c2833d088.png)
![patched_font_10](https://user-images.githubusercontent.com/10674169/58262412-31d36180-7db5-11e9-98b7-4cea9bd68d07.png)

## Contributor 🎊
- [mattn](https://github.com/mattn)
- [yasu47b](https://github.com/yasu47b)
- [b4b4r07](https://github.com/b4b4r07)
- [delphinus](https://github.com/delphinus)
- [TsutomuNakamura](https://github.com/TsutomuNakamura)

## Author
[codehex](https://twitter.com/CodeHex)

