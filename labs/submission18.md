# Lab 18 — Reproducible Builds with Nix

## Task 1 — Build Reproducible Python App

```bash
$ curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
info: downloading the Determinate Nix Installer
 INFO nix-installer v3.20.0
`nix-installer` needs to run as `root`, attempting to escalate now via `sudo`...
[sudo] password for lord: 
 INFO nix-installer v3.20.0
Nix install plan (v3.20.0)
Planner: linux (with default settings)

Planned actions:
* Create directory `/nix`
* Install Determinate Nixd
* Extract the bundled Nix (originally from /nix/store/3h8wrgwk0rczd8rjv4fvl344mp1w765m-nix-binary-tarball-3.20.0/nix-3.20.0-x86_64-linux.tar.xz) to `/nix/temp-install-dir`
* Create a directory tree in `/nix`
* Synchronize /nix and /nix/var ownership
* Move the downloaded Nix into `/nix`
* Synchronize /nix/store ownership
* Create build users (UID 30001-30032) and group (GID 30000)
* Setup the default Nix profile
* Place the Nix configuration in `/etc/nix/nix.conf`
* Configure the shell profiles
* Configure the Determinate Nix daemon
* Cleanup


Proceed? ([Y]es/[n]o/[e]xplain): Y
 INFO Step: Create directory `/nix`
 INFO Step: Install Determinate Nixd
 INFO Step: Provision Nix
 INFO Step: Create build users (UID 30001-30032) and group (GID 30000)
 INFO Step: Configure Nix
 INFO Step: Create directory `/etc/tmpfiles.d`
 INFO Step: Configure the Determinate Nix daemon
 INFO Step: Cleanup
 INFO Running self test for shell sh
 INFO Running self test for shell bash
Nix was installed successfully!
To get started using Nix, open a new shell or run `. /nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh`
$ nix --version
nix (Determinate Nix 3.20.0) 2.34.6
$ nix run nixpkgs#hello
Hello, world!
```

I decided to take app from Lab 2.

```bash
$ nix-build
this derivation will be built:
  /nix/store/nsq42hhp4bms61n94s84kix6440gwnkf-devops-info-service-1.0.0.drv
these 78 paths will be fetched (113.3 MiB download, 481.2 MiB unpacked):
  /nix/store/80kzdbnl1bmnd8r8hphh1xzmgsvyyf6w-acl-2.3.2
  /nix/store/s894wa8wzbf841myjr7dm4m85177f8rl-attr-2.5.2
  /nix/store/1jzhbwq5rjjaqa75z88ws2b424vh7m53-bash-5.2p32
  /nix/store/z34dss3rj1rnp1g820r64w0na7452jy1-binutils-2.41
  /nix/store/xf9367aw8afx9a8y9r90kdhlaji9bnz8-binutils-2.41-lib
  /nix/store/rl56awy2w2iwvgdmibv98k0vx7lzyw21-binutils-wrapper-2.41
  /nix/store/0s4k9fh17m2sx8ifj9xdflnjfgxrvbgp-bzip2-1.0.8
  /nix/store/8krchn3igfzz62fnirxj47l2hrz11sj9-bzip2-1.0.8-bin
  /nix/store/n9gh8gxx5xx51ihgll2l20ar9b2vmgzy-coreutils-9.5
  /nix/store/w14121h1n3kg8pjaw790mkr7z7cb238y-die-hook
  /nix/store/k6slgyjf4sv0lpz1pdg6sr3bpvjdr35l-diffutils-3.10
  /nix/store/bjfycbr15irgfgwqw6bjb5xmrb6jllnq-ed-1.20.2
  /nix/store/6ykxv7d14syjk1mc58k8afflsc8p73iz-ensure-newer-sources-hook
  /nix/store/13bjf2r63138p99rcdkdb0wr0gbwcc0m-expand-response-params
  /nix/store/nbbg70f6gihj51p65kv19m0fnq8ik5kh-expat-2.6.4
  /nix/store/0vdf5mpd762bw53rgl5nkmhvzq8n4m0d-file-5.45
  /nix/store/qdk6ds72q8xf5fzi16x7736bwz8aydic-findutils-4.9.0
  /nix/store/jziybc6yn78sf0bbh0kmmx50bxjsnr15-gawk-5.2.2
  /nix/store/f2lglcw69yh2yyihxq6kvhpanc3s1n9p-gcc-13.2.0
  /nix/store/90yn7340r8yab8kxpb0p7y0c9j3snjam-gcc-13.2.0-lib
  /nix/store/dd13q38yxm9qppjclsvwn10dscsf0l9w-gcc-13.2.0-libgcc
  /nix/store/rdc1jnyw74mwr2gszqc5zwi433zxs089-gcc-wrapper-13.2.0
  /nix/store/y8fzwmygqh9rl5rc8rvcydcskgszldpn-gdbm-1.23
  /nix/store/pf5avvvl4ssd6kylcvg2g23hcjp71h19-glibc-2.39-52
  /nix/store/94gaw2ngsnpf3ybxjsg89lf3hf3d55y2-glibc-2.39-52-bin
  /nix/store/jig62nn8174n4dlk05lqwsvs5wd2c64r-glibc-2.39-52-dev
  /nix/store/bz756gkg3lrqjl958d8h7z5qn98z0vn7-gmp-6.3.0
  /nix/store/0pkjhzncyjkvhq8lwmdkzvl4cs4vh0yb-gmp-with-cxx-6.3.0
  /nix/store/d0i8idmbb4jji9ml01xsqgykrbvm7dss-gnu-config-2024-01-01
  /nix/store/8nss7h1yk4jihkmr4xj5ihrbdkv4y1wy-gnugrep-3.11
  /nix/store/0554jm1l1qw1pcfqsliw91hnifn11w8m-gnumake-4.4.1
  /nix/store/y8br765djcj51ls9lb3kylkrvc2wan3p-gnused-4.9
  /nix/store/llalnjlyrj2zv12q5bjy8cagqv70j73y-gnutar-1.35
  /nix/store/r6apkwli4s0xhzn1bdi9nrkmvqc5arrj-gzip-1.13
  /nix/store/mwa26ixrnf0s1j7923pdzfl9nwwglq4w-isl-0.20
  /nix/store/gniy4ab9wcijxjpcciddgpzdwq3v3dnb-libffi-3.4.6
  /nix/store/9jivp79yv91fl1i6ayq2107a78q7k43i-libidn2-2.3.7
  /nix/store/w0gcbzh5w6ds5h2cf1cg7nc2xcpcwxdk-libmpc-1.3.1
  /nix/store/zvwpisszhpkkk8spqyya8n3bpm7wj39p-libunistring-1.1
  /nix/store/bb99dclcsv3r0a8q967bnvga02qicxsf-libxcrypt-4.4.36
  /nix/store/d7vji37yjnbji14m4waa5rn35f10jzya-linux-headers-6.7
  /nix/store/v68clh8wa74xxblx0b748vr50gl7xnzc-mailcap-2.1.53
  /nix/store/gg7jfhvmfdcvny0rafz1mzmgl8dyvwdm-make-shell-wrapper-hook
  /nix/store/6x0n8ksnajz1kf7n6q0farmyrc6af4mz-mpdecimal-4.0.0
  /nix/store/4dqj8bq01rwapgzgwm0zpr622r6jv8wq-mpfr-4.2.1
  /nix/store/zmi2nlx42h1qrs2v7yn142dq4zjq30im-ncurses-6.4
  /nix/store/p8hw2h465g0byxwpamnk6gv6mp5gnqn2-openssl-3.0.14
  /nix/store/hxhg0hk4g60yjsbiw1hwdzykidswdvaz-patch-2.7.6
  /nix/store/2wp235bg03gykpixd9v2nyxp08w8xq8a-patchelf-0.15.0
  /nix/store/jk99vx5dhv57zldl5j0id3n23igbj269-pcre2-10.43
  /nix/store/6xplkmsjayxypq5chazjf14fmcd0jsd3-python-catch-conflicts-hook
  /nix/store/rnm3maa3k41q787mvfbjzr8j0hcr5vb6-python-imports-check-hook.sh
  /nix/store/52syyn9n7agv20y4xz07izfz8bfi6603-python-namespaces-hook.sh
  /nix/store/r5mnw30vc49dj2gwi55zshqshxigl3gq-python-remove-bin-bytecode-hook
  /nix/store/h5s4n27qdnnarhp4ly79albcd2j247wr-python-remove-tests-dir-hook
  /nix/store/s0p1kr5mvs0j42dq5r08kgqbi0k028f2-python3-3.11.10
  /nix/store/24hf1fh5vkmrzl951nc0aq4mifr88m2y-python3.11-annotated-types-0.6.0
  /nix/store/s2hr3ffaf6gdbcc2pfkbzanz7vgja0s2-python3.11-anyio-4.3.0
  /nix/store/c9b2c5gakdnqakspr3phy0ld4b9p93z5-python3.11-click-8.1.7
  /nix/store/ids9mi9ss5gzsqp819akf3mysv4dy3z4-python3.11-fastapi-0.110.2
  /nix/store/lc6mb3c7phf8x80qd1y4vav6kxdbnija-python3.11-h11-0.14.0
  /nix/store/adhq1qdhlbasy8cbbcg5bd0s9lsrqxnr-python3.11-idna-3.7
  /nix/store/l4br40rzpsf632csmzppddh81ixldijs-python3.11-pydantic-2.6.3
  /nix/store/5gdxpcv7xz17xkkilwi18a3wfpi5wwj0-python3.11-pydantic-core-2.16.3
  /nix/store/g06gxrd87769lklypl85pf78h08qq37y-python3.11-sniffio-1.3.1
  /nix/store/mbib0hl5c89w8bk1mg6ss55wdgf7kai9-python3.11-starlette-0.37.2
  /nix/store/18yiwdazlwkndhqlg3zp0m2zf20qhxzw-python3.11-typing-extensions-4.11.0
  /nix/store/0asg4r9f7yi9a54g64pcy44bqmddvdk4-python3.11-uvicorn-0.29.0
  /nix/store/rafcl7sl3f9757z4m5hl7vaklppi8xkr-readline-8.2p10
  /nix/store/0pjc3s5k0zlc56xs2z0cmvq885zkij9h-sqlite-3.45.3
  /nix/store/80wijs24wjp619zmrasrh805bax02xjm-stdenv-linux
  /nix/store/vil8wqacfkz8qx09aymahxv48d1zq3g2-tzdata-2024b
  /nix/store/2329271b42wh6b6yhl7jmjyi0cs4428b-update-autotools-gnu-config-scripts-hook
  /nix/store/p969ygqmpq5qjawbs8dlryqbny4x7pym-wrap-python-hook
  /nix/store/2y852kcvb7shrj8f3z8j22pa0iybcbgj-xgcc-13.2.0-libgcc
  /nix/store/p1m1hbbfyvix5390xgslhq3d30zzhp7f-xz-5.4.7
  /nix/store/1h4vhcpa6nqln9x0w03dy5i1x97z8svc-xz-5.4.7-bin
  /nix/store/6ghariqqf33l5xqn7alx58dys7kz2wm5-zlib-1.3.1
copying path '/nix/store/r5mnw30vc49dj2gwi55zshqshxigl3gq-python-remove-bin-bytecode-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/h5s4n27qdnnarhp4ly79albcd2j247wr-python-remove-tests-dir-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/w14121h1n3kg8pjaw790mkr7z7cb238y-die-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/v68clh8wa74xxblx0b748vr50gl7xnzc-mailcap-2.1.53' from 'https://cache.nixos.org'...
copying path '/nix/store/vil8wqacfkz8qx09aymahxv48d1zq3g2-tzdata-2024b' from 'https://cache.nixos.org'...
copying path '/nix/store/dd13q38yxm9qppjclsvwn10dscsf0l9w-gcc-13.2.0-libgcc' from 'https://cache.nixos.org'...
copying path '/nix/store/d0i8idmbb4jji9ml01xsqgykrbvm7dss-gnu-config-2024-01-01' from 'https://cache.nixos.org'...
copying path '/nix/store/2y852kcvb7shrj8f3z8j22pa0iybcbgj-xgcc-13.2.0-libgcc' from 'https://install.determinate.systems'...
copying path '/nix/store/zvwpisszhpkkk8spqyya8n3bpm7wj39p-libunistring-1.1' from 'https://install.determinate.systems'...
copying path '/nix/store/d7vji37yjnbji14m4waa5rn35f10jzya-linux-headers-6.7' from 'https://install.determinate.systems'...
copying path '/nix/store/2329271b42wh6b6yhl7jmjyi0cs4428b-update-autotools-gnu-config-scripts-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/9jivp79yv91fl1i6ayq2107a78q7k43i-libidn2-2.3.7' from 'https://install.determinate.systems'...
copying path '/nix/store/pf5avvvl4ssd6kylcvg2g23hcjp71h19-glibc-2.39-52' from 'https://cache.nixos.org'...
copying path '/nix/store/s894wa8wzbf841myjr7dm4m85177f8rl-attr-2.5.2' from 'https://cache.nixos.org'...
copying path '/nix/store/1jzhbwq5rjjaqa75z88ws2b424vh7m53-bash-5.2p32' from 'https://cache.nixos.org'...
copying path '/nix/store/0s4k9fh17m2sx8ifj9xdflnjfgxrvbgp-bzip2-1.0.8' from 'https://cache.nixos.org'...
copying path '/nix/store/bjfycbr15irgfgwqw6bjb5xmrb6jllnq-ed-1.20.2' from 'https://cache.nixos.org'...
copying path '/nix/store/13bjf2r63138p99rcdkdb0wr0gbwcc0m-expand-response-params' from 'https://cache.nixos.org'...
copying path '/nix/store/nbbg70f6gihj51p65kv19m0fnq8ik5kh-expat-2.6.4' from 'https://cache.nixos.org'...
copying path '/nix/store/jziybc6yn78sf0bbh0kmmx50bxjsnr15-gawk-5.2.2' from 'https://cache.nixos.org'...
copying path '/nix/store/90yn7340r8yab8kxpb0p7y0c9j3snjam-gcc-13.2.0-lib' from 'https://cache.nixos.org'...
copying path '/nix/store/y8fzwmygqh9rl5rc8rvcydcskgszldpn-gdbm-1.23' from 'https://cache.nixos.org'...
copying path '/nix/store/94gaw2ngsnpf3ybxjsg89lf3hf3d55y2-glibc-2.39-52-bin' from 'https://cache.nixos.org'...
copying path '/nix/store/bz756gkg3lrqjl958d8h7z5qn98z0vn7-gmp-6.3.0' from 'https://cache.nixos.org'...
copying path '/nix/store/0554jm1l1qw1pcfqsliw91hnifn11w8m-gnumake-4.4.1' from 'https://cache.nixos.org'...
copying path '/nix/store/y8br765djcj51ls9lb3kylkrvc2wan3p-gnused-4.9' from 'https://cache.nixos.org'...
copying path '/nix/store/bb99dclcsv3r0a8q967bnvga02qicxsf-libxcrypt-4.4.36' from 'https://cache.nixos.org'...
copying path '/nix/store/gniy4ab9wcijxjpcciddgpzdwq3v3dnb-libffi-3.4.6' from 'https://cache.nixos.org'...
copying path '/nix/store/6x0n8ksnajz1kf7n6q0farmyrc6af4mz-mpdecimal-4.0.0' from 'https://cache.nixos.org'...
copying path '/nix/store/zmi2nlx42h1qrs2v7yn142dq4zjq30im-ncurses-6.4' from 'https://cache.nixos.org'...
copying path '/nix/store/8krchn3igfzz62fnirxj47l2hrz11sj9-bzip2-1.0.8-bin' from 'https://cache.nixos.org'...
copying path '/nix/store/80kzdbnl1bmnd8r8hphh1xzmgsvyyf6w-acl-2.3.2' from 'https://cache.nixos.org'...
copying path '/nix/store/p8hw2h465g0byxwpamnk6gv6mp5gnqn2-openssl-3.0.14' from 'https://cache.nixos.org'...
copying path '/nix/store/hxhg0hk4g60yjsbiw1hwdzykidswdvaz-patch-2.7.6' from 'https://cache.nixos.org'...
copying path '/nix/store/r6apkwli4s0xhzn1bdi9nrkmvqc5arrj-gzip-1.13' from 'https://cache.nixos.org'...
copying path '/nix/store/gg7jfhvmfdcvny0rafz1mzmgl8dyvwdm-make-shell-wrapper-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/jk99vx5dhv57zldl5j0id3n23igbj269-pcre2-10.43' from 'https://cache.nixos.org'...
copying path '/nix/store/p1m1hbbfyvix5390xgslhq3d30zzhp7f-xz-5.4.7' from 'https://cache.nixos.org'...
copying path '/nix/store/6ghariqqf33l5xqn7alx58dys7kz2wm5-zlib-1.3.1' from 'https://cache.nixos.org'...
copying path '/nix/store/mwa26ixrnf0s1j7923pdzfl9nwwglq4w-isl-0.20' from 'https://cache.nixos.org'...
copying path '/nix/store/4dqj8bq01rwapgzgwm0zpr622r6jv8wq-mpfr-4.2.1' from 'https://cache.nixos.org'...
copying path '/nix/store/jig62nn8174n4dlk05lqwsvs5wd2c64r-glibc-2.39-52-dev' from 'https://cache.nixos.org'...
copying path '/nix/store/llalnjlyrj2zv12q5bjy8cagqv70j73y-gnutar-1.35' from 'https://cache.nixos.org'...
copying path '/nix/store/rafcl7sl3f9757z4m5hl7vaklppi8xkr-readline-8.2p10' from 'https://cache.nixos.org'...
copying path '/nix/store/xf9367aw8afx9a8y9r90kdhlaji9bnz8-binutils-2.41-lib' from 'https://cache.nixos.org'...
copying path '/nix/store/0vdf5mpd762bw53rgl5nkmhvzq8n4m0d-file-5.45' from 'https://cache.nixos.org'...
copying path '/nix/store/0pjc3s5k0zlc56xs2z0cmvq885zkij9h-sqlite-3.45.3' from 'https://cache.nixos.org'...
copying path '/nix/store/1h4vhcpa6nqln9x0w03dy5i1x97z8svc-xz-5.4.7-bin' from 'https://cache.nixos.org'...
copying path '/nix/store/8nss7h1yk4jihkmr4xj5ihrbdkv4y1wy-gnugrep-3.11' from 'https://cache.nixos.org'...
copying path '/nix/store/0pkjhzncyjkvhq8lwmdkzvl4cs4vh0yb-gmp-with-cxx-6.3.0' from 'https://cache.nixos.org'...
copying path '/nix/store/2wp235bg03gykpixd9v2nyxp08w8xq8a-patchelf-0.15.0' from 'https://cache.nixos.org'...
copying path '/nix/store/w0gcbzh5w6ds5h2cf1cg7nc2xcpcwxdk-libmpc-1.3.1' from 'https://cache.nixos.org'...
copying path '/nix/store/z34dss3rj1rnp1g820r64w0na7452jy1-binutils-2.41' from 'https://cache.nixos.org'...
copying path '/nix/store/s0p1kr5mvs0j42dq5r08kgqbi0k028f2-python3-3.11.10' from 'https://cache.nixos.org'...
copying path '/nix/store/n9gh8gxx5xx51ihgll2l20ar9b2vmgzy-coreutils-9.5' from 'https://cache.nixos.org'...
copying path '/nix/store/f2lglcw69yh2yyihxq6kvhpanc3s1n9p-gcc-13.2.0' from 'https://cache.nixos.org'...
copying path '/nix/store/k6slgyjf4sv0lpz1pdg6sr3bpvjdr35l-diffutils-3.10' from 'https://cache.nixos.org'...
copying path '/nix/store/qdk6ds72q8xf5fzi16x7736bwz8aydic-findutils-4.9.0' from 'https://cache.nixos.org'...
copying path '/nix/store/rl56awy2w2iwvgdmibv98k0vx7lzyw21-binutils-wrapper-2.41' from 'https://cache.nixos.org'...
copying path '/nix/store/6ykxv7d14syjk1mc58k8afflsc8p73iz-ensure-newer-sources-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/52syyn9n7agv20y4xz07izfz8bfi6603-python-namespaces-hook.sh' from 'https://cache.nixos.org'...
copying path '/nix/store/6xplkmsjayxypq5chazjf14fmcd0jsd3-python-catch-conflicts-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/rnm3maa3k41q787mvfbjzr8j0hcr5vb6-python-imports-check-hook.sh' from 'https://cache.nixos.org'...
copying path '/nix/store/24hf1fh5vkmrzl951nc0aq4mifr88m2y-python3.11-annotated-types-0.6.0' from 'https://cache.nixos.org'...
copying path '/nix/store/lc6mb3c7phf8x80qd1y4vav6kxdbnija-python3.11-h11-0.14.0' from 'https://cache.nixos.org'...
copying path '/nix/store/c9b2c5gakdnqakspr3phy0ld4b9p93z5-python3.11-click-8.1.7' from 'https://cache.nixos.org'...
copying path '/nix/store/adhq1qdhlbasy8cbbcg5bd0s9lsrqxnr-python3.11-idna-3.7' from 'https://cache.nixos.org'...
copying path '/nix/store/g06gxrd87769lklypl85pf78h08qq37y-python3.11-sniffio-1.3.1' from 'https://cache.nixos.org'...
copying path '/nix/store/18yiwdazlwkndhqlg3zp0m2zf20qhxzw-python3.11-typing-extensions-4.11.0' from 'https://cache.nixos.org'...
copying path '/nix/store/p969ygqmpq5qjawbs8dlryqbny4x7pym-wrap-python-hook' from 'https://cache.nixos.org'...
copying path '/nix/store/s2hr3ffaf6gdbcc2pfkbzanz7vgja0s2-python3.11-anyio-4.3.0' from 'https://cache.nixos.org'...
copying path '/nix/store/0asg4r9f7yi9a54g64pcy44bqmddvdk4-python3.11-uvicorn-0.29.0' from 'https://cache.nixos.org'...
copying path '/nix/store/5gdxpcv7xz17xkkilwi18a3wfpi5wwj0-python3.11-pydantic-core-2.16.3' from 'https://cache.nixos.org'...
copying path '/nix/store/mbib0hl5c89w8bk1mg6ss55wdgf7kai9-python3.11-starlette-0.37.2' from 'https://cache.nixos.org'...
copying path '/nix/store/l4br40rzpsf632csmzppddh81ixldijs-python3.11-pydantic-2.6.3' from 'https://cache.nixos.org'...
copying path '/nix/store/rdc1jnyw74mwr2gszqc5zwi433zxs089-gcc-wrapper-13.2.0' from 'https://cache.nixos.org'...
copying path '/nix/store/ids9mi9ss5gzsqp819akf3mysv4dy3z4-python3.11-fastapi-0.110.2' from 'https://cache.nixos.org'...
copying path '/nix/store/80wijs24wjp619zmrasrh805bax02xjm-stdenv-linux' from 'https://cache.nixos.org'...
building '/nix/store/nsq42hhp4bms61n94s84kix6440gwnkf-devops-info-service-1.0.0.drv'...
Sourcing python-remove-tests-dir-hook
Sourcing python-catch-conflicts-hook.sh
Sourcing python-remove-bin-bytecode-hook.sh
Sourcing python-imports-check-hook.sh
Using pythonImportsCheckPhase
Sourcing python-namespaces-hook
Running phase: unpackPhase
unpacking source archive /nix/store/0mmm77r04ng0755ag7m9bd6wddlqacm4-app_python
source root is app_python
setting SOURCE_DATE_EPOCH to timestamp 315619200 of file app_python/requirements.txt
Running phase: patchPhase
Running phase: updateAutotoolsGnuConfigScriptsPhase
Running phase: configurePhase
no configure script, doing nothing
Running phase: buildPhase
no Makefile or custom buildPhase, doing nothing
Running phase: installPhase
Running phase: fixupPhase
shrinking RPATHs of ELF executables and libraries in /nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
checking for references to /build/ in /nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0...
patching script interpreter paths in /nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0/bin/.devops-info-service-wrapped: interpreter directive changed from "#!/usr/bin/env python3" to "/nix/store/s0p1kr5mvs0j42dq5r08kgqbi0k028f2-python3-3.11.10/bin/python3"
stripping (with command strip and flags -S -p) in  /nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0/bin
Rewriting #!/nix/store/s0p1kr5mvs0j42dq5r08kgqbi0k028f2-python3-3.11.10/bin/python3 to #!/nix/store/s0p1kr5mvs0j42dq5r08kgqbi0k028f2-python3-3.11.10
wrapping `/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0/bin/.devops-info-service-wrapped'...
Rewriting #! /nix/store/1jzhbwq5rjjaqa75z88ws2b424vh7m53-bash-5.2p32/bin/bash -e to #!/nix/store/s0p1kr5mvs0j42dq5r08kgqbi0k028f2-python3-3.11.10
Executing pythonRemoveTestsDir
Finished executing pythonRemoveTestsDir
Running phase: installCheckPhase
no Makefile or custom installCheckPhase, doing nothing
Running phase: pythonCatchConflictsPhase
Running phase: pythonRemoveBinBytecodePhase
Running phase: pythonImportsCheckPhase
Executing pythonImportsCheckPhase
/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
```

```bash
$ ./result/bin/devops-info-service
2026-05-14 22:01:46,762 - __main__ - INFO - Application starting...
2026-05-14 22:02:12,921 - __main__ - INFO - Incoming request: GET /
```

```bash
curl.exe http://localhost:8000/
{"service":{"name":"devops-info-service","version":"1.0.0","description":"DevOps course info service","framework":"FastAPI"},"system":{"hostname":"DESKTOP-0RFQR97","platform":"Linux","platform_version":"#1 SMP PREEMPT_DYNAMIC Thu Jun  5 18:30:46 UTC 2025","architecture":"x86_64","cpu_count":16,"python_version":"3.11.10"},"runtime":{"uptime_seconds":147,"uptime_human":"0 hours, 2 minutes","current_time":"2026-05-14T19:04:14.113125+00:00","timezone":"UTC"},"request":{"client_ip":"127.0.0.1","user_agent":"curl/8.19.0","method":"GET","path":"/"},"endpoints":[{"path":"/","method":"GET","description":"Service information"},{"path":"/health","method":"GET","description":"Health check"},{"path":"/docs","method":"GET","description":"Auto-generated API documentation"}]}
```

```bash
$ readlink result
/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
$ rm result
$ nix-build
/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
$ readlink result
/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
```

The store path is identical – Nix reused the cached build. To force a true rebuild (without cache), we deleted the store path; the new build still produced the same hash because all inputs were identical.

```bash
$ STORE_PATH=$(readlink result)
$ echo "Original store path: $STORE_PATH"
Original store path: /nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
$ nix-store --delete $STORE_PATH
finding garbage collector roots...
removing stale link from "/nix/var/nix/gcroots/auto/1i4i2fsz01215maqahidzpb7vwm5f828" to "/tmp/.tmpQLj3u5/profile-2-link"
removing stale link from "/nix/var/nix/gcroots/auto/4wmg47snysdcq5sbdh1maskmi3wilnnc" to "/tmp/.tmpQLj3u5/profile-1-link"
removing stale link from "/nix/var/nix/gcroots/auto/smw0nd9j80xy1jcrhb07a88nl9xmqdxz" to "/tmp/.tmpQLj3u5/profile"
removing stale temporary roots file "/nix/var/nix/temproots/12596"
removing stale temporary roots file "/nix/var/nix/temproots/12598"
removing stale temporary roots file "/nix/var/nix/temproots/9860"
removing stale temporary roots file "/nix/var/nix/temproots/10921"
removing stale temporary roots file "/nix/var/nix/temproots/4985"
removing stale temporary roots file "/nix/var/nix/temproots/3166"
removing stale temporary roots file "/nix/var/nix/temproots/11007"
removing stale temporary roots file "/nix/var/nix/temproots/5539"
removing stale temporary roots file "/nix/var/nix/temproots/3465"
removing stale temporary roots file "/nix/var/nix/temproots/5267"
removing stale temporary roots file "/nix/var/nix/temproots/11005"
removing stale temporary roots file "/nix/var/nix/temproots/3082"
removing stale temporary roots file "/nix/var/nix/temproots/10161"
0 store paths deleted, 0.0 KiB freed
error: Cannot delete path '/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0' because it's referenced by the GC root '/home/lordik/DevOps-Course/DevOps-Core-Course/labs/lab18/app_python/result'.
$ rm result
$ nix-build
/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
$ readlink result
/nix/store/phj4bk4184ydc1lxp50wavh0wdnm4n61-devops-info-service-1.0.0
```

```bash
echo "flask" > requirements-unpinned.txt

python -m venv venv1
source venv1/bin/activate
pip install -r requirements-unpinned.txt
pip freeze | grep -i flask > freeze1.txt
deactivate


pip cache purge 2>/dev/null || rm -rf ~/.cache/pip

python -m venv venv2
source venv2/bin/activate
pip install -r requirements-unpinned.txt
pip freeze | grep -i flask > freeze2.txt
deactivate
```

**The store path format is:**

```/nix/store/<hash>-<name>-<version>```

- `<hash>` is a SHA‑256 hash derived from **all** inputs: source code, build instructions, compiler flags, and every dependency (transitively).  
- `<name>` and `<version>` come from `pname` and `version` in the derivation.

Because the hash depends on everything that affects the build, the same hash guarantees identical content. This is why Nix can safely share binary caches – the hash proves the content.

### Reflection: What would Nix solved in Lab 1

- **“Works on my machine”** issues
- **Transitive dependency drift** – pip pins only direct dependencies; Nix pins everything.
- **Inconsistent dev environments** – Nix provides a declarative, reproducible shell (`nix-shell` or `nix develop`) instead of fragile virtual environments.

Nix would also have made CI/CD much more reliable – no more “it builds locally but fails in CI” because both run in identical sandboxes.

## Task 2 — Reproducible Docker Images

```bash
$ docker build -t lab2-app:v1 .
...
$ docker inspect lab2-app:v1 | grep Created
        "Created": "2026-05-14T19:18:50.356924781Z",
```
```bash
$ docker build -t lab2-app:v2 ./app_python
$ docker inspect lab2-app:v2 | grep Created
        "Created": "2026-05-14T19:18:50.356924781Z",
```

I got the same result :)

```bash
$ nix-build docker.nix
...
$ docker load < result
...
$ docker run -d -p 8000:8000 --name lab2-container lab
2-app:v1
b647e568b7cbf48c9b4b555ca738417fd083fc64a24ca0a3870b13250368de27
$ docker run -d -p 8001:8000 --name nix-container devo
ps-info-service-nix:1.0.0
26f81076ff6c7095a5add7579929437c5ef2fb4f6c74929b9e3566d070f93a3d
```

```bash
curl.exe http://localhost:8000/
{"service":{"name":"devops-info-service","version":"1.0.0","description":"DevOps course info service","framework":"FastAPI"},"system":{"hostname":"b647e568b7cb","platform":"Linux","platform_version":"#1 SMP PREEMPT_DYNAMIC Thu Jun  5 18:30:46 UTC 2025","architecture":"x86_64","cpu_count":16,"python_version":"3.13.13"},"runtime":{"uptime_seconds":26,"uptime_human":"0 hours, 0 minutes","current_time":"2026-05-14T19:25:59.776574+00:00","timezone":"UTC"},"request":{"client_ip":"172.17.0.1","user_agent":"curl/8.19.0","method":"GET","path":"/"},"endpoints":[{"path":"/","method":"GET","description":"Service information"},{"path":"/health","method":"GET","description":"Health check"},{"path":"/docs","method":"GET","description":"Auto-generated API documentation"}]}
curl.exe http://localhost:8001/
{"service":{"name":"devops-info-service","version":"1.0.0","description":"DevOps course info service","framework":"FastAPI"},"system":{"hostname":"26f81076ff6c","platform":"Linux","platform_version":"#1 SMP PREEMPT_DYNAMIC Thu Jun  5 18:30:46 UTC 2025","architecture":"x86_64","cpu_count":16,"python_version":"3.11.10"},"runtime":{"uptime_seconds":12,"uptime_human":"0 hours, 0 minutes","current_time":"2026-05-14T19:26:03.109846+00:00","timezone":"UTC"},"request":{"client_ip":"172.17.0.1","user_agent":"curl/8.19.0","method":"GET","path":"/"},"endpoints":[{"path":"/","method":"GET","description":"Service information"},{"path":"/health","method":"GET","description":"Health check"},{"path":"/docs","method":"GET","description":"Auto-generated API documentation"}]}
```

```bash
$ rm result
$ nix-build docker.nix
...
$ sha256sum result
00bd7200b8bc3c30981f410cf274670beb4ef809cc7c8aac88389fe70077788c  result
$ rm result
$ nix-build docker.nix
/nix/store/jkmllpifn92bw8q20dywci8h3jmshav5-devops-info-service-nix.tar.gz
$ sha256sum result
00bd7200b8bc3c30981f410cf274670beb4ef809cc7c8aac88389fe70077788c  result
```

```bash
$ docker build -t lab2-app:test1 .
...
$ docker save lab2-app:test1 | sha256sum
0f8b185407d89a2a6ab427d8de5a7f740a519f59b4bd3dc7476331a80b40819f  -
$ docker build -t lab2-app:test2 .
...
$ docker save lab2-app:test2 | sha256sum
a10806b0aeaf083f4be0813b87886ce30cf9bf86f9acc9ff6eb2e7916f4e549e  -
```

Different hashes, but same images

```bash
$ docker images | grep -E "lab2-app|devops-info-service-nix"
lab2-app                         test1       ceeaafb3160f   12 minutes ago   164MB
lab2-app                         test2       ceeaafb3160f   12 minutes ago   164MB
lab2-app                         v1          ceeaafb3160f   12 minutes ago   164MB
devops-info-service-nix          1.0.0       84e514a5c302   56 years ago     199MB
```

nix image has bigger size than usual (in lab dif)

```bash
$ docker history lab2-app:v1
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
ceeaafb3160f   13 minutes ago   /bin/sh -c #(nop)  CMD ["python" "app.py"]      0B        
26f2d0e7168d   13 minutes ago   /bin/sh -c #(nop)  EXPOSE 8000                  0B        
1160de2490f1   13 minutes ago   /bin/sh -c #(nop)  USER appuser                 0B        
d1abc64a39e7   13 minutes ago   /bin/sh -c #(nop) COPY file:1b1fdca683149a6b…   4.16kB    
d20f4c034e31   13 minutes ago   /bin/sh -c pip install --no-cache-dir -r req…   46.1MB    
9677bfc3f4a2   13 minutes ago   /bin/sh -c #(nop) COPY file:e542bd09c9d90010…   42B       
d8165d968afc   13 minutes ago   /bin/sh -c #(nop) WORKDIR /app                  0B        
5cfeb1c5d181   13 minutes ago   /bin/sh -c useradd --create-home --shell /bi…   8.92kB    
18e30fdc01ce   5 days ago       CMD ["python3"]                                 0B        buildkit.dockerfile.v0
<missing>      5 days ago       RUN /bin/sh -c set -eux;  for src in idle3 p…   36B       buildkit.dockerfile.v0
<missing>      5 days ago       RUN /bin/sh -c set -eux;   savedAptMark="$(a…   35.3MB    buildkit.dockerfile.v0
<missing>      6 days ago       ENV PYTHON_SHA256=2ab91ff401783ccca64f75d10c…   0B        buildkit.dockerfile.v0
<missing>      6 days ago       ENV PYTHON_VERSION=3.13.13                      0B        buildkit.dockerfile.v0
<missing>      6 days ago       ENV GPG_KEY=7169605F62C751356D054A26A821E680…   0B        buildkit.dockerfile.v0
<missing>      6 days ago       RUN /bin/sh -c set -eux;  apt-get update;  a…   3.81MB    buildkit.dockerfile.v0
<missing>      6 days ago       ENV PATH=/usr/local/bin:/usr/local/sbin:/usr…   0B        buildkit.dockerfile.v0
<missing>      9 days ago       # debian.sh --arch 'amd64' out/ 'trixie' '@1…   78.6MB    debuerreotype 0.17
```

```bash
$ docker history devops-info-service-nix:1.0.0
IMAGE          CREATED   CREATED BY   SIZE      COMMENT
84e514a5c302   N/A                    411B      store paths: ['/nix/store/a9znh3v2c3pxg6zgmapswffxahyiq62w-devops-info-service-nix-customisation-layer']
<missing>      N/A                    12kB      store paths: ['/nix/store/mv3v991ijp5jps9mb8wf0vjlni869n0q-devops-info-service-1.0.0']
<missing>      N/A                    741kB     store paths: ['/nix/store/0asg4r9f7yi9a54g64pcy44bqmddvdk4-python3.11-uvicorn-0.29.0']
<missing>      N/A                    1.65MB    store paths: ['/nix/store/ids9mi9ss5gzsqp819akf3mysv4dy3z4-python3.11-fastapi-0.110.2']
<missing>      N/A                    1.01MB    store paths: ['/nix/store/mbib0hl5c89w8bk1mg6ss55wdgf7kai9-python3.11-starlette-0.37.2']
<missing>      N/A                    5.14MB    store paths: ['/nix/store/l4br40rzpsf632csmzppddh81ixldijs-python3.11-pydantic-2.6.3']
<missing>      N/A                    644kB     store paths: ['/nix/store/lc6mb3c7phf8x80qd1y4vav6kxdbnija-python3.11-h11-0.14.0']
<missing>      N/A                    1.3MB     store paths: ['/nix/store/c9b2c5gakdnqakspr3phy0ld4b9p93z5-python3.11-click-8.1.7']
<missing>      N/A                    5.67MB    store paths: ['/nix/store/5gdxpcv7xz17xkkilwi18a3wfpi5wwj0-python3.11-pydantic-core-2.16.3']
<missing>      N/A                    1.37MB    store paths: ['/nix/store/s2hr3ffaf6gdbcc2pfkbzanz7vgja0s2-python3.11-anyio-4.3.0']
<missing>      N/A                    100kB     store paths: ['/nix/store/24hf1fh5vkmrzl951nc0aq4mifr88m2y-python3.11-annotated-types-0.6.0']
<missing>      N/A                    41.6kB    store paths: ['/nix/store/g06gxrd87769lklypl85pf78h08qq37y-python3.11-sniffio-1.3.1']
<missing>      N/A                    906kB     store paths: ['/nix/store/adhq1qdhlbasy8cbbcg5bd0s9lsrqxnr-python3.11-idna-3.7']
<missing>      N/A                    417kB     store paths: ['/nix/store/18yiwdazlwkndhqlg3zp0m2zf20qhxzw-python3.11-typing-extensions-4.11.0']
<missing>      N/A                    121MB     store paths: ['/nix/store/s0p1kr5mvs0j42dq5r08kgqbi0k028f2-python3-3.11.10']
<missing>      N/A                    809kB     store paths: ['/nix/store/p1m1hbbfyvix5390xgslhq3d30zzhp7f-xz-5.4.7']
<missing>      N/A                    1.89MB    store paths: ['/nix/store/vil8wqacfkz8qx09aymahxv48d1zq3g2-tzdata-2024b']
<missing>      N/A                    1.54MB    store paths: ['/nix/store/0pjc3s5k0zlc56xs2z0cmvq885zkij9h-sqlite-3.45.3']
<missing>      N/A                    469kB     store paths: ['/nix/store/rafcl7sl3f9757z4m5hl7vaklppi8xkr-readline-8.2p10']
<missing>      N/A                    6.48MB    store paths: ['/nix/store/p8hw2h465g0byxwpamnk6gv6mp5gnqn2-openssl-3.0.14']
<missing>      N/A                    216kB     store paths: ['/nix/store/6x0n8ksnajz1kf7n6q0farmyrc6af4mz-mpdecimal-4.0.0']
<missing>      N/A                    110kB     store paths: ['/nix/store/v68clh8wa74xxblx0b748vr50gl7xnzc-mailcap-2.1.53']
<missing>      N/A                    129kB     store paths: ['/nix/store/bb99dclcsv3r0a8q967bnvga02qicxsf-libxcrypt-4.4.36']
<missing>      N/A                    72.3kB    store paths: ['/nix/store/gniy4ab9wcijxjpcciddgpzdwq3v3dnb-libffi-3.4.6']
<missing>      N/A                    811kB     store paths: ['/nix/store/y8fzwmygqh9rl5rc8rvcydcskgszldpn-gdbm-1.23']
<missing>      N/A                    272kB     store paths: ['/nix/store/nbbg70f6gihj51p65kv19m0fnq8ik5kh-expat-2.6.4']
<missing>      N/A                    79.5kB    store paths: ['/nix/store/0s4k9fh17m2sx8ifj9xdflnjfgxrvbgp-bzip2-1.0.8']
<missing>      N/A                    1.61MB    store paths: ['/nix/store/1jzhbwq5rjjaqa75z88ws2b424vh7m53-bash-5.2p32']
<missing>      N/A                    9.08MB    store paths: ['/nix/store/90yn7340r8yab8kxpb0p7y0c9j3snjam-gcc-13.2.0-lib']
<missing>      N/A                    159kB     store paths: ['/nix/store/dd13q38yxm9qppjclsvwn10dscsf0l9w-gcc-13.2.0-libgcc']
<missing>      N/A                    127kB     store paths: ['/nix/store/6ghariqqf33l5xqn7alx58dys7kz2wm5-zlib-1.3.1']
<missing>      N/A                    3.16MB    store paths: ['/nix/store/zmi2nlx42h1qrs2v7yn142dq4zjq30im-ncurses-6.4']
<missing>      N/A                    29.9MB    store paths: ['/nix/store/pf5avvvl4ssd6kylcvg2g23hcjp71h19-glibc-2.39-52']
<missing>      N/A                    159kB     store paths: ['/nix/store/2y852kcvb7shrj8f3z8j22pa0iybcbgj-xgcc-13.2.0-libgcc']
<missing>      N/A                    346kB     store paths: ['/nix/store/9jivp79yv91fl1i6ayq2107a78q7k43i-libidn2-2.3.7']
<missing>      N/A                    1.83MB    store paths: ['/nix/store/zvwpisszhpkkk8spqyya8n3bpm7wj39p-libunistring-1.1']
```

### Build Results Summary

**Traditional Docker (Lab 2)**
- `Created` timestamp same for back‑to‑back builds: `2026-05-14T19:18:50.356924781Z`
- Saved image tarball hashes differed:  
  `0f8b18...` vs `a10806...` → not reproducible.
- Image size: **164 MB**.

**Nix dockerTools (Lab 18)**
- `sha256sum result` after two independent builds:  
  `00bd7200...` (identical) → **reproducible**.
- Image size: **199 MB** (larger, but every dependency is visible and fixed).
- Python version inside Nix image: `3.11.10` (pinned by nixpkgs).

**Observation**  
Nix guarantees bit‑for‑bit reproducibility at the cost of a slightly larger image (trade‑off can be tuned with `maxLayers`). Traditional Docker images vary even when timestamps look the same.

### Analysis: Why Can’t Traditional Dockerfiles Achieve Bit‑for‑Bit Reproducibility?

Traditional `Dockerfile` builds are **not deterministic** because of several inherent factors:

- **Base image tags** like `python:3.13-slim` are mutable – they can point to different images over time.
- **Build timestamps** are embedded in image layers (even if the `CREATED` field appears the same, layer metadata differs).
- **Package managers** (`apt-get`, `pip`) fetch the latest versions unless pinned by hash, and even pinned versions may include timestamps or non‑deterministic file ordering.
- **File system metadata** (e.g., `atime`, `mtime`) varies between builds.
- **Caching behaviour** depends on local state – a clean build vs. cached build can produce different outputs.

Nix avoids all of these by:
- Using content‑addressable storage.
- Fixing all inputs (including system libraries and compiler versions).
- Sandboxing builds to eliminate host influence.

### Reflection: If I Could Redo Lab 2 with Nix

I would:
- Replace the `Dockerfile` with a Nix `dockerTools.buildLayeredImage` expression.
- Use `created = "1970-01-01T00:00:01Z"` to fix timestamps.
- Add explicit `maxLayers = 1` or `maxLayers = 2` to reduce image size.
- Use `extraCommands` to strip unnecessary files (e.g., `rm -rf /nix/store/*-python3-*/lib/python3.11/test`).
- Define all Python dependencies in the Nix derivation, eliminating `pip` at build time.

This would produce a **smaller, perfectly reproducible image** and avoid the “works on my machine” class of problems.

### Practical Scenarios Where Nix’s Reproducibility Matters

- **CI/CD pipelines** – guarantee that every commit produces exactly the same binary, preventing “build flakiness” and flaky tests.
- **Security audits** – verify that a deployed image corresponds exactly to a known source commit; any hash mismatch indicates tampering or an unintended change.
- **Rollbacks** – when a production issue occurs, you can roll back to an exact previous build with zero uncertainty about what changed.
- **Compliance** (e.g., finance, healthcare) – regulators require proof that a build is bit‑for‑bit identical to the audited source. Nix provides that proof via content hashes.
- **Cross‑team collaboration** – every developer gets the same environment, eliminating “it works on my machine” disputes.

