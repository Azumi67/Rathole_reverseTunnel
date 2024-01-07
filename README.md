![R (2)](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/a064577c-9302-4f43-b3bf-3d4f84245a6f)
نام پروژه : ریورس تانل Rathole [ با TCP - UDP - WS + TLS ]
---------------------------------------------------------------

-----------
![Update-Note--Arvin61r58](https://github.com/Azumi67/Rathole_reverseTunnel/assets/119934376/ae94b74a-2870-4d17-b886-ebfaadf7cc0c)**اپدیت**

- تغییراتی در دستورات compile انجام شد.
- تغییراتی در math اسکریپت انجام شد.
- متود دوم برای سرورهای ایرانی که مشکل دریافت self cert داشتند اضافه شد.
- اموزش نوشتاری نصب و compile به صورت دستی هم اضافه شد
- اگر مشکلی در دانلود داشتید، میتوانید از dns های شکن و غیره استفاده کنید.
- لطفا کانکشن بدون هدر استفاده کنید.
---------------

**توضیح کوتاه در مورد این پروژه :**
- **اگر سرعتتون پایین بود، لطفا هم بر روی سرور ایران و خارج optimizer نصب کنید.**
- اسکریپت بارها تست شده و باگ هایش گرفته شده است.
- حتما در سرور تست، نخست تانل را ازمایش کنید و سپس اقدام به استفاده از آن بکنید.
- تمامی تست های من با سرورهای کاملا فیلتر شده بوده است.
- در این اسکریپت شما میتوانید 10 سرور خارج را به یک سرور ایران وصل کنید ولی در تانل Ws + TLS تنها 5 سرور خارج را به یک سرور ایران، میتوانید وصل نمایید.
- حتما اگر در تانل به مشکلی خوردید،لاگ سرویس را بررسی نمایید. systemctl status kharej-azumi و systemctl status iran-azumi
- از TCP و UDP پشتیبانی میکند.
- ریست تایمر برای سرویس های خود را بر اساس نیاز خودتان تعیین کنید.
- حتما ریست تایمر سرور خارج و ایران یکسان باشد.
- اگر از این تانل راضی بودید، میشه بعدا 2 سرور ایران هم اضافه کرد.
- حتما در صورت مشکل دانلود، dns های خود را تغییر دهید.
- قبل از اجرای اسکریپت اصلی ، باید اسکریپت نصب پروژه را اجرا کنید.(اگر خطایی در compile کردن پروژه گرفتید، در اینترنت سرچ کنید و مشکل compile نشدن پروژه را بیابید)
- پنل شما در خارج باید نصب شده باشد
- اگر به هر دلیلی پیش نیاز ها برای شما نصب نشد و خطا گرفتید، لطفا با قرار دادن DNS دوباره امتحان بفرمایید.
- اگر اختلالی در تانل داشتید همیشه وارد مسیر روبرو شوید cd /etc/systemd/system و با دستور ls ، سرویس های خارج و ایران را بیابید و با دستور systemctl status servicename و یا journalctl -u servicename.service ، دلیل اختلال تانل را بیابید

------------------------

![307981](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/39b2794b-fd04-4ae5-baea-d4b66138766e)
 **فهرست :**
 ----------------------
 - **دسترسی سریع به اسکریپت** >> **[کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%A7%D8%B3%DA%A9%D8%B1%DB%8C%D9%BE%D8%AA-%D9%85%D9%86)**
- **نصب go به صورت دستی یا با اسکریپت** >> [کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%A7%D9%85%D9%88%D8%B2%D8%B4-%D9%86%D8%B5%D8%A8-go-%D9%85%D9%88%D8%B1%D8%AF-%D9%86%DB%8C%D8%A7%D8%B2-%D8%A8%D8%B1%D8%A7%DB%8C-%D8%A7%D8%AC%D8%B1%D8%A7%DB%8C-%D8%A7%D8%B3%DA%A9%D8%B1%DB%8C%D9%BE%D8%AA)
- **نصب پروژه** >> [کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D9%86%D8%B5%D8%A8-%D9%BE%D8%B1%D9%88%DA%98%D9%87)
- **امکانات** >> [کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%A7%D9%85%DA%A9%D8%A7%D9%86%D8%A7%D8%AA)
- **تانل ریورس TCP ایپی 4**  >> **[کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%AA%D8%A7%D9%86%D9%84-%D8%B1%DB%8C%D9%88%D8%B1%D8%B3-tcp-%D8%A7%DB%8C%D9%BE%DB%8C-4)**
- **تانل ریورس TCP ایپی 6** >> **[کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%AA%D8%A7%D9%86%D9%84-%D8%B1%DB%8C%D9%88%D8%B1%D8%B3-tcp-%D8%A7%DB%8C%D9%BE%DB%8C-6)
- **تانل ریورس UDP ایپی 4**  >> **[کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%AA%D8%A7%D9%86%D9%84-%D8%B1%DB%8C%D9%88%D8%B1%D8%B3-udp-%D8%A7%DB%8C%D9%BE%DB%8C-4)
- **تانل ریورس UDP ایپی 6**  >> **[کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%AA%D8%A7%D9%86%D9%84-%D8%B1%DB%8C%D9%88%D8%B1%D8%B3-udp-%D8%A7%DB%8C%D9%BE%DB%8C-6)
- **تانل ریورس TLS + WS ایپی 4** >> [کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%AA%D8%A7%D9%86%D9%84-%D8%B1%DB%8C%D9%88%D8%B1%D8%B3-tls--ws-%D8%A7%DB%8C%D9%BE%DB%8C-4)
- **تانل ریورس TLS + WS ایپی 6** >> [کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%AA%D8%A7%D9%86%D9%84-%D8%B1%DB%8C%D9%88%D8%B1%D8%B3-tls--ws-%D8%A7%DB%8C%D9%BE%DB%8C-6)
- **اسکریپت های کارآمد** >> [کلیک](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%A7%D8%B3%DA%A9%D8%B1%DB%8C%D9%BE%D8%AA-%D9%87%D8%A7%DB%8C-%DA%A9%D8%A7%D8%B1%D8%A2%D9%85%D8%AF-)

--------------
![4961639](https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/ac38de6c-00d6-47e8-a455-e0201143db1c)**نصب پروژه**
-

- با این اسکریپت [click](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%A7%D8%B3%DA%A9%D8%B1%DB%8C%D9%BE%D8%AA-%D9%85%D9%86) ، نخست پروژه را بر روی تمامی سرور ها نصب نمایید. من تنها بر روی amd64 و سیستم عامل های دبیان و اوبونتو تست کردم و دسترسی به سایر را نداشتم.
- توجه فرمایید که نردیک به 2 گیگ فضای خالی برای compile نیاز دارید.
- وقتی اسکریپت اجرا شد گزینه یک را انتخاب کنید تا rust برای شما نصب شود.
- اگر خطای مبنی بر [profile.release] گرفتید، داخل مسیر nano rathole/Cargo.toml شوید و زیر [profile.release] این strip = true را به strip = "symbols" تغییر دهید اگر خطا در این باب بود.
- هر خطایی در compile پروژه گرفتید در داخل اینترنت سرچ نمایید و مشکل خود را حل نمایید.
- میتوانید بر روی یک سرور تازه ریبلد شده تست بفرمایید و حتما قبلش سرور را اپدیت کرده باشید و dns هم تنظیم کنید.
- **نصب به صورت دستی**
```
sudo apt update -y
apt install rustc -y
apt install cargo -y
apt-get install pkg-config libssl-dev -y
apt install curl -y
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
choose 1 
sudo apt install git -y
git clone https://github.com/miyugundam/rathole.git
if you got an error by [workspace] , add it here nano rathole/Cargo.toml
cd rathole
cargo build

```
- اگر خطای روبرو را گرفتید
```
Caused by:
  failed to parse the edition key
Caused by:
  supported edition values are 2015 or 2018, but 2021 is unknown
```

 - دوباره اقدام به نصب از طریق اسکریپت زیر نمایید
  ```
  sudo apt install curl -y && bash <(curl -s https://raw.githubusercontent.com/Azumi67/Rathole_reverseTunnel/main/install3.sh)
  ```
-----------------


![Go_(programming_language)-Logo wine](https://github.com/Azumi67/UDP2RAW_FEC/assets/119934376/71b80a34-9515-42de-8238-9065986104a1)**اموزش نصب go مورد نیاز برای اجرای اسکریپت**
-
- شما میتوانید از طریق اسکریپت [Here](https://github.com/Azumi67/Rathole_reverseTunnel#%D8%A7%D8%B3%DA%A9%D8%B1%DB%8C%D9%BE%D8%AA-%D9%85%D9%86) ، این پیش نیاز را نصب کنید یا به صورت دستی نصب نمایید.
- حتما در صورت مشکل دانلود، dns های خود را تغییر دهید.
- پس از نصب پیش نیاز ، اجرای اسکریپت go برای بار اول، ممکن است تا 10 ثانیه طول بکشد اما بعد از آن سریع اجرا میشود.
- یا نصب به صورت دستی :
```
sudo apt update
arm64 : wget https://go.dev/dl/go1.21.5.linux-arm64.tar.gz
arm64 : sudo tar -C /usr/local -xzf go1.21.5.linux-arm64.tar.gz

amd64 : wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
amd64 : sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz

nano ~/.bash_profile
paste this into it : export PATH=$PATH:/usr/local/go/bin
save and exit with Ctrl + x , then Y

source ~/.bash_profile
go mod init mymodule
go mod tidy
go get github.com/AlecAivazis/survey/v2
go get github.com/fatih/color
go get github.com/pkg/sftp
go get -u golang.org/x/crypto/ssh

```
- سپس اسکریپت را میتوانید اجرا نمایید.
------------------
![check](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/13de8d36-dcfe-498b-9d99-440049c0cf14)
**امکانات**
-

- پشتیبانی از TCP و UDP
- قابلیت تانل بر روی چندین پورت
- امکان استفاده از ایپی 4 و 6
- ریست تایمر انتخابی توسط شما
- مناسب برای v2ray , Wireguard
- امکان تانل بر روی ایپی فیلتر شده
- امکان استفاده از ریورس تانل Ws + TLS
- امکان تانل بر روی 10 سرور خارج و یک سرور ایران
- امکان استفاده از nodelay برای بهبود پینگ
- امکان ترکیب ریورس تانل udp با fec در آینده 
- امکان حذف تمامی تانل ها و سرویس ها

 ------------------------------------------------------

![147-1472495_no-requirements-icon-vector-graphics-clipart](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/98d8c2bd-c9d2-4ecf-8db9-246b90e1ef0f)
 **پیش نیازها**

 - لطفا سرور اپدیت شده باشه.
 - میتوانید از اسکریپت اقای [Hwashemi](https://github.com/hawshemi/Linux-Optimizer) و یا [OPIRAN](https://github.com/opiran-club/VPS-Optimizer) هم برای بهینه سازی سرور در صورت تمایل استفاده نمایید. (پیش نیاز نیست)


----------------------------

  
  ![6348248](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/398f8b07-65be-472e-9821-631f7b70f783)
**آموزش**
-
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل ریورس TCP ایپی 4** 
----------------------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران**

**مسیر : IPV4 TCP > IRAN**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/f742316d-550f-4cce-81ed-b14739be19fd" alt="Image" />
</p>



- نخست سرور ایران را کانفیگ میکنم
- قبلش باید پروژه را بر روی سرورهای خود نصب و compile نمایید.
- من دو سرور خارج دارم و هر سرور خارج من، دو کانفیگ دارد.
- در سرور ایران از من میخواهد تعداد کل کانفیگ هایم را وارد کنم. من دو سرور خارج دارم و هر سرور دو عدد کانفیگ دارد، پس باید عدد 4 را وارد کنم. 
- پورت تانل را 443 وارد میکنم.
- پورت های سرور های خارجم را به ترتیب وارد میکنم. دقت نمایید در این قسمت باید تمامی پورت های سرور خارج را وارد نمایید
- به طور مثال اگر در سرور اول خارج، دو کانفیگ با پورت های 8080 و 8081 و در سرور دوم خارج، دو کانفیگ دیگر با پورت های 8082 و 8083 دارم. پس به ترتیب، تمام این پورت ها را وارد میکنم.
- اگر پینگ پایین تری میخواهید، nodelay را در ازای کاهش پهنای باند فعال نمایید.
- ریست تایمر را هم هر 2 ساعت انتخاب میکنم.
----------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج اول** 

**مسیر : IPV4 TCP > KHAREJ 1**



 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/95cc5720-547d-4d4f-80d3-eb70ee448c30" alt="Image" />
</p>

- سرور اول خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number برای سرور اول خارج، همیشه عدد یک میباشد.
- ایپی 4 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 2 عدد کانفیگ در سرور خارج داشتم. پس عدد 2 را وارد میکنم.
- پورت های کانفیگ سرور اول خارج 8080 و 8081 بود.
- گزینه nodelay هم که در سرور ایران فعال کرده بودیم پس در اینجا هم فعال میکنم.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.


--------------------------------------

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج دوم** 

**مسیر : IPV4 TCP > KHAREJ 2**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/725f7669-4bfc-4ef1-b386-34ecdfef4b37" alt="Image" />
</p>

- سرور دوم خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number در سرور اول به ما نمایش داده شد که عدد 3 بود پس عدد 3، برای سرور دوم خارج را وارد میکنیم.
- ایپی 4 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 2 عدد کانفیگ در سرور خارج دارم. پس عدد 2 را وارد میکنم.
- پورت های کانفیگ سرور دوم خارج 8082 و 8083 بود.
- گزیینه nodelay هم که در سرور ایران فعال کرده بودیم پس در اینجا هم فعال میکنم.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.

--------------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل ریورس TCP ایپی 6** 
-
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران**

**مسیر : IPV6 TCP > IRAN**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/afbc5bb6-b371-438b-bfd6-32c0fb503269" alt="Image" />
</p>


- نخست سرور ایران را کانفیگ میکنم
- قبلش باید پروژه را بر روی سرورهای خود نصب و compile نماییم.
- من دو سرور خارج دارم و هر سرور خارج من دو کانفیگ دارد.
- در سرور ایران از من میخواهد تعداد کل کانفیگ هایم را وارد کنم. من دو سرور خارج داشتم و هر سرور دو عدد کانفیگ دارد، پس باید عدد 4 را وارد کنم. 
- پورت تانل را 443 وارد میکنم.
- پورت های سرور های خارجم را به ترتیب وارد میکنم. دقت نمایید در این قسمت باید تمامی پورت های سرور خارج را وارد نمایید
- به طور مثال اگر در سرور اول خارج، دو کانفیگ با پورت های 8080 و 8081 و در سرور دوم خارج، دو کانفیگ دیگر با پورت های 8082 و 8083 دارم. پس به ترتیب، تمام این پورت ها را وارد میکنم.
- اگر پینگ پایین تری میخواهید، nodelay را در ازای کاهش پهنای باند فعال نمایید.
- ریست تایمر را هم هر 2 ساعت انتخاب میکنم.
----------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج اول** 

**مسیر : IPV6 TCP > KHAREJ 1**



 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/8a7c6eb0-7db1-427e-9b78-54a1853aa72e" alt="Image" />
</p>

- سرور اول خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number برای سرور اول خارج، همیشه عدد یک میباشد.
- ایپی 6 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 2 عدد کانفیگ در سرور خارج داشتم. پس عدد 2 را وارد میکنم.
- پورت های کانفیگ سرور اول خارج 8080 و 8081 بود.
- گزیینه nodelay هم که در سرور ایران فعال کرده بودیم پس در اینجا هم فعال میکنم.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.


--------------------------------------

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج دوم** 

**مسیر : IPV6 TCP > KHAREJ 2**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/23b7b97b-1584-422b-8af6-c66ff1af54be" alt="Image" />
</p>

- سرور دوم خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number در سرور اول به ما نمایش داده شد که عدد 3 بود پس عدد 3، برای سرور دوم خارج را وارد میکنیم.
- ایپی 6 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 2 عدد کانفیگ در سرور خارج دارم. پس عدد 2 را وارد میکنم.
- پورت های کانفیگ سرور دوم خارج 8082 و 8083 بود.
- گزیینه nodelay هم که در سرور ایران فعال کرده بودیم پس در اینجا هم فعال میکنم.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.



--------------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل ریورس UDP ایپی 4** 
-
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران**

**مسیر : IPV4 UDP > IRAN**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/1eebab65-a6d3-4eef-a2bc-099b54fd1db5" alt="Image" />
</p>


- نخست سرور ایران را کانفیگ میکنم
- قبلش باید پروژه را بر روی سرورهای خود نصب و compile نماییم.
- من یک سرور خارج دارم و سرور خارج من یک کانفیگ دارد.
- در سرور ایران از من میخواهد تعداد کل کانفیگ هایم را وارد کنم. من یک سرور خارج داشتم و سرور من یک عدد کانفیگ دارد، پس باید عدد 1 را وارد کنم. 
- پورت تانل را 443 وارد میکنم.
- پورت وایرگارد سرور خارجم را وارد میکنم. دقت نمایید در این قسمت باید تمامی پورت های سرور خارج را وارد نمایید
- چون من یک کانفیگ وایرگارد دارم، پس تنها یک پورت را وارد میکنم. پورت من 50820 میباشد
- ریست تایمر را هم هر 2 ساعت انتخاب میکنم.
----------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج اول** 

**مسیر : IPV4 UDP > KHAREJ 1**



 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/b9e3576d-fdf2-413d-9d54-16f5cd2c5cd6" alt="Image" />
</p>

- سرور خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number برای سرور اول خارج، همیشه عدد یک میباشد.
- ایپی 4 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 1 عدد کانفیگ در سرور خارج داشتم. پس عدد 1 را وارد میکنم.
- پورت کانفیگ سرور  خارج 50820 بود.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.

--------------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل ریورس UDP ایپی 6** 
-
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران**

**مسیر : IPV6 UDP > IRAN**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/d72ddc20-b7b6-4426-b5ca-3cb5136a808d" alt="Image" />
</p>


- نخست سرور ایران را کانفیگ میکنم
- قبلش باید پروژه را بر روی سرورهای خود نصب و compile نماییم.
- من یک سرور خارج دارم و سرور خارج من یک کانفیگ دارد.
- در سرور ایران از من میخواهد تعداد کل کانفیگ هایم را وارد کنم. من یک سرور خارج داشتم و سرور من یک عدد کانفیگ دارد، پس باید عدد 1 را وارد کنم. 
- پورت تانل را 443 وارد میکنم.
- پورت وایرگارد سرور خارجم را وارد میکنم. دقت نمایید در این قسمت باید تمامی پورت های سرور خارج را وارد نمایید
- چون من یک کانفیگ وایرگارد دارم، پس تنها یک پورت را وارد میکنم. پورت من 50820 میباشد
- ریست تایمر را هم هر 2 ساعت انتخاب میکنم.
----------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج اول** 

**مسیر : IPV6 UDP > KHAREJ 1**



 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/f6b49336-f3e5-4fc6-8ce3-6713fd33b0b9" alt="Image" />
</p>

- سرور خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number برای سرور اول خارج، همیشه عدد یک میباشد.
- ایپی 6 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 1 عدد کانفیگ در سرور خارج داشتم. پس عدد 1 را وارد میکنم.
- پورت کانفیگ سرور  خارج 50820 بود.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.
----------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل ریورس TLS + WS ایپی 4** 
-
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران**

**مسیر : IPV4 WS + TLS > IRAN**

 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/b761011a-1401-43fa-87c7-cad23c160051" alt="Image" />
</p>

- از منو، گزینه اول را انتخاب میکنیم تا سرور ایران را کانفیگ نماییم و cert لازمه رو برای سرور دریافت نماییم.باید rootCA.crt را در تمامی سرور های خارج در پوشه rathole پیست کنیم. شما میتوانید یا از طریق copy cert اینکار از طریق اسکریپت انجام دهید یا خودتان به صورت دستی ، rootCA.crt را در سرور خارج کپی کنید.
- اگر اینکار را نکنید ، ارتباط برقرار نخواهد شد.


<p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/5155020d-86bc-490d-860b-b9ee75cead3c" alt="Image" />
</p>

- سرور ایران را کانفیگ میکنیم
- قبلش باید پروژه را بر روی سرورهای خود نصب و compile نماییم.
- من دو سرور خارج دارم و هر سرور خارج من دو کانفیگ دارد.
- در سرور ایران از من میخواهد تعداد کل کانفیگ هایم را وارد کنم. من دو سرور خارج داشتم و هر سرور دو عدد کانفیگ دارد، پس باید عدد 4 را وارد کنم. 
- پورت تانل را 443 وارد میکنم.
- پورت های سرور های خارجم را به ترتیب وارد میکنم. دقت نمایید در این قسمت باید تمامی پورت های سرور خارج را وارد نمایید
- به طور مثال اگر در سرور اول خارج، دو کانفیگ با پورت های 8080 و 8081 و در سرور دوم خارج، دو کانفیگ دیگر با پورت های 8082 و 8083 دارم. پس به ترتیب، تمام این پورت ها را وارد میکنم.
- ریست تایمر را هم هر 2 ساعت انتخاب میکنم.


<p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/e222732b-463a-461c-8da9-50b186cbb7f8" alt="Image" />
</p>

- خب پس از کانفیگ ایران، باید rootCA.crt را در تمامی سرورهای خارجمان، کپی کنیم. من در اینجا بوسیله اسکریپت اینکار را انجام میدم.
- دقت نمایید که باید بتوانید به صورت دستی هم از سرور ایران به سرور خارج، ssh بزنید در غیر اینصورت با اسکریپت هم امکان پذیر نخواهد بود.
- ایپی 4 خارج و پورت ssh سرور خارج هم وارد میکنم.
- یوزرنیم و پسورد سرور خارج هم وارد میکنم و فایل rootCA.crt به صورت اتوماتیک به پوشه مورد نظر در سرور خارج انتقال داده میشود.
- حتما قبل از کانفیگ، اطمینان پیدا کنید که در تمامی سرور های خارج و ایران شما، پروژه نصب شده باشد
----------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج اول** 

**مسیر : IPV4 WS + TLS > Kharej 1**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/783aec92-506c-4c72-b5a0-e76e9b851753" alt="Image" />
</p>

- سرور اول خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number برای سرور اول خارج، همیشه عدد یک میباشد.
- ایپی 4 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 2 عدد کانفیگ در سرور خارج داشتم. پس عدد 2 را وارد میکنم.
- پورت های کانفیگ سرور اول خارج 8080 و 8081 بود.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.


--------------------------------------

![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج دوم** 

**مسیر : IPV4 WS + TLS > Kharej 2**


 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/98cd4d22-3664-47f6-9f07-6d39813a77ee" alt="Image" />
</p>

- سرور دوم خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number در سرور اول به ما نمایش داده شد که عدد 3 بود پس عدد 3، برای سرور دوم خارج را وارد میکنیم.
- ایپی 4 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 2 عدد کانفیگ در سرور خارج دارم. پس عدد 2 را وارد میکنم.
- پورت های کانفیگ سرور دوم خارج 8082 و 8083 بود.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.
--------------------------------------
![OIP2 (1)](https://github.com/Azumi67/V2ray_loadbalance_multipleServers/assets/119934376/3ec2f05f-3308-4441-8cce-62ab4776f4e2)
**تانل ریورس TLS + WS ایپی 6** 
-
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور ایران**

**مسیر : IPV6 WS + TLS > IRAN**

 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/b761011a-1401-43fa-87c7-cad23c160051" alt="Image" />
</p>

- از منو، گزینه اول را انتخاب میکنیم تا سرور ایران را کانفیگ نماییم و cert لازمه رو برای سرور دریافت نماییم.باید rootCA.crt را در تمامی سرور های خارج در پوشه rathole پیست کنیم. شما میتوانید یا از طریق copy cert اینکار از طریق اسکریپت انجام دهید یا خودتان به صورت دستی ، rootCA.crt را در سرور خارج کپی کنید.
- اگر اینکار را نکنید ، ارتباط برقرار نخواهد شد.
<p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/044bf137-6195-4716-81bf-a54ed2b8f3fc" alt="Image" />
</p>

- سرور ایران را کانفیگ میکنیم
- قبلش باید پروژه را بر روی سرورهای خود نصب و compile نماییم.
- من یک سرور خارج دارم و سرور خارج من دو کانفیگ دارد.
- در سرور ایران از من میخواهد تعداد کل کانفیگ هایم را وارد کنم. من یک سرور خارج داشتم و سرور من دو عدد کانفیگ دارد، پس باید عدد 2 را وارد کنم. 
- پورت تانل را 443 وارد میکنم.
- پورت های سرور های خارجم را به ترتیب وارد میکنم. دقت نمایید در این قسمت باید تمامی پورت های سرور خارج را وارد نمایید
- پورت های من 8080 و 8081 میباشد.
- ریست تایمر را هم هر 2 ساعت انتخاب میکنم.

<p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/e222732b-463a-461c-8da9-50b186cbb7f8" alt="Image" />
</p>

- خب پس از کانفیگ ایران، باید rootCA.crt را در سرور خارج، کپی کنیم. من در اینجا بوسیله اسکریپت اینکار را انجام میدم.
- دقت نمایید که باید بتوانید به صورت دستی هم از سرور ایران به سرور خارج، ssh بزنید در غیر اینصورت با اسکریپت هم امکان پذیر نخواهد بود.
- ایپی 4 خارج و پورت ssh سرور خارج هم وارد میکنم.
- یوزرنیم و پسورد سرور خارج هم وارد میکنم و فایل rootCA.crt به صورت اتوماتیک به پوشه مورد نظر در سرور خارج انتقال داده میشود.
- حتما قبل از کانفیگ، اطمینان پیدا کنید که در تمامی سرور های خارج و ایران شما، پروژه نصب شده باشد
----------------------
![green-dot-clipart-3](https://github.com/Azumi67/6TO4-PrivateIP/assets/119934376/902a2efa-f48f-4048-bc2a-5be12143bef3) **سرور خارج اول** 

**مسیر : IPV6 WS + TLS > Kharej 1**

 <p align="right">
  <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/17832681-a6e3-436b-b61e-82f6936c8b20" alt="Image" />
</p>

- سرور خارج را کانفیگ میکنم.
- از انجا که این ریورس تانل شبیه frp میباشد، من هم از starting number برای جدا کردن کانفیگ ها استفاده خواهم کرد.
- مقدار starting number برای سرور اول خارج، همیشه عدد یک میباشد.
- ایپی 6 ایران را وارد میکنم.
- پورت تانل که 443 قرار داده بودم
- من 2 عدد کانفیگ در سرور خارج داشتم. پس عدد 2 را وارد میکنم.
- پورت های کانفیگ سرور خارج 8080 و 8081 بود.
- ریست تایمر هم که عدد 2 را وارد کرده بودیم. ( باید ریست تایمر یکسان باشد که همه سرویس ها همزمان ریست شوند)
- در اخر یک عدد به شما نشان داده میشود. در سرور خارج بعدی وقتی از شما مقدار starting number را خواست، عددی که به شما نمایش داده شده است را وارد نمایید.


--------------------------------------

**اسکرین شات**

<details>
  <summary align="right">Click to reveal image</summary>
  
  <p align="right">
    <img src="https://github.com/Azumi67/Fast_reverseTunnel/assets/119934376/eb2a6a2b-0628-41de-9ca3-1d442a05ee24" alt="menu screen" />
  </p>
</details>


------------------------------------------
![scri](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/cbfb72ac-eff1-46df-b5e5-a3930a4a6651)
**اسکریپت های کارآمد :**
-
- این اسکریپت ها optional میباشد.


 
 Opiran Scripts
 
```
 bash <(curl -s https://raw.githubusercontent.com/opiran-club/pf-tun/main/pf-tun.sh --ipv4)
```

```
apt install curl -y && bash <(curl -s https://raw.githubusercontent.com/opiran-club/VPS-Optimizer/main/optimizer.sh --ipv4)
```

Hawshemi script

```
wget "https://raw.githubusercontent.com/hawshemi/Linux-Optimizer/main/linux-optimizer.sh" -O linux-optimizer.sh && chmod +x linux-optimizer.sh && bash linux-optimizer.sh
```

-----------------------------------------------------
![R (a2)](https://github.com/Azumi67/PrivateIP-Tunnel/assets/119934376/716fd45e-635c-4796-b8cf-856024e5b2b2)
**اسکریپت من**
----------------
**کامپایل پروژه**
- دستور زیر پروژه را بر روی سرور شما compile میکند. نخست این را اجرا نمایید
```
sudo apt install curl -y && bash <(curl -s https://raw.githubusercontent.com/Azumi67/Rathole_reverseTunnel/main/install.sh)

```
- اگر با نصب git مشکل داشتید اول git را جداگانه نصب کنید و سپس اقدام به اجرای اسکریپت کنید
```
sudo apt update -y
sudo apt --fix-broken install 
sudo apt install git -y
nano ~/.bashrc
paste this into it  >>  export PATH="$PATH:/usr/bin/git"
Ctrl + x and enter y
source ~/.bashrc
git --version
```
- حالا اقدام به اجرای اسکریپت نمایید
```
sudo apt install curl -y && bash <(curl -s https://raw.githubusercontent.com/Azumi67/Rathole_reverseTunnel/main/install2.sh)
```

- اگر خطای روبرو را گرفتید

```
Caused by:
  failed to parse the edition key
Caused by:
  supported edition values are 2015 or 2018, but 2021 is unknown
 ```

  - دوباره اقدام به نصب از طریق اسکریپت زیر نمایید
  ```
  sudo apt install curl -y && bash <(curl -s https://raw.githubusercontent.com/Azumi67/Rathole_reverseTunnel/main/install3.sh)
  ```

------------------
  
**اسکرپت تانل**
- پس از انکه کامپایل پروژه تمام شد، با دستور زیر، پیش نیاز گو را نصب میکنید و سپس اقدام به اجرای اسکریپت میکنید.
```
sudo apt install curl -y && bash <(curl -s https://raw.githubusercontent.com/Azumi67/Rathole_reverseTunnel/main/go.sh)
```

- اگر به صورت دستی پیش نیاز های گو را نصب کردید و میخواهید به صورت دستی هم اسکریپت را اجرا کنید میتوانید با دستور زیر انجام دهید
```
sudo apt install wget -y && rm rat.go && wget https://raw.githubusercontent.com/Azumi67/Rathole_reverseTunnel/main/rat.go && go run rat.go
```

---------------------------------------------
![R23 (1)](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/18d12405-d354-48ac-9084-fff98d61d91c)
**سورس ها**


![R (9)](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/33388f7b-f1ab-4847-9e9b-e8b39d75deaa) [سورس  RAThole](https://github.com/rapiz1/rathole)

![R (9)](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/33388f7b-f1ab-4847-9e9b-e8b39d75deaa) [سورس  OPIRAN](https://github.com/opiran-club)

![R (9)](https://github.com/Azumi67/6TO4-GRE-IPIP-SIT/assets/119934376/4758a7da-ab54-4a0a-a5a6-5f895092f527)[سورس  Hwashemi](https://github.com/hawshemi/Linux-Optimizer)



-----------------------------------------------------

![youtube-131994968075841675](https://github.com/Azumi67/FRP-V2ray-Loadbalance/assets/119934376/24202a92-aff2-4079-a6c2-9db14cd0ecd1)
**ویدیوی آموزش**

-----------------------------------------
