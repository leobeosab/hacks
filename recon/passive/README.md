Shitty passive recon tool written in golang. Waiting for github.com/leoebosab/sharingan to be completed

## TODO
* Embeds for DiscordMessage (stop showing links)
* Out-of-scope black list
* Rate limiting for Gobuster
* Notify flag for on even non unique data for testing
* Notify on unique dirb results
* Ability to set Webhooks from scan 

* nmap scan
* waybackurls
* aquatone
* what web


## The Steps
1.  Read targets -done
2.  DNS bust required targets using gobuster dns and amass enum -d -done
    *   If subs.txt exists compare with subs.txt -done
    *   Keep list of new subs -done
    *   Output file to scan json -done
3.  Nmap scan all new subs
    *   add nmap scan results to struct?
    *   note any cool ports
5.  What web
5.  probed subs get dirbd -done
6.  waybackurls
    
