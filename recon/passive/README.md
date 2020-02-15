Shitty passive recon tool written in golang. Waiting for github.com/leoebosab/sharingan to be completed

## TODO
1. Multiple Webhooks
2. Context for commands
3. Embeds for DiscordMessage (stop showing links)
4. Out-of-scope black list
5. Rate limiting for Gobuster
6. Notify on even non unique data for testing
7. Notify on unique dirb results

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
    
