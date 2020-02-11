## TODO
* dns scan wild cards
* save
* get differences
* nmap scan
* dirbust
* aquatone
* discord messages
* what web


## The Steps
1.  Read targets
    *   If target has * that means dns bust
2.  DNS bust required targets using gobuster dns and amass enum -d
    *   If subs.txt exists compare with subs.txt
    *   Keep list of new subs
    *   Output file to found subs.txt
3.  Nmap scan all new subs
    *   add nmap scan results to struct?
    *   note any cool ports
4.  httprobe new subs
5.  What web
5.  probed subs get dirbd
6.  waybackurls
7.  Aquatone?
    
