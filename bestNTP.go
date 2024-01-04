package main

import (
	"fmt"
	"sort"
	"time"
	"github.com/beevik/ntp" // https://github.com/beevik/ntp
)

type ServerInfo struct {
	Server         string
	RTT            time.Duration
	RootDispersion time.Duration
}

func main() {
	ntpServers := []string{
		// Global NTP Servers
		"0.beevik-ntp.pool.ntp.org",
		"1.beevik-ntp.pool.ntp.org",
		"2.beevik-ntp.pool.ntp.org",

		// German NTP Servers
		"0.de.pool.ntp.org",
		"1.de.pool.ntp.org",
		"2.de.pool.ntp.org",
		"3.de.pool.ntp.org",

		// European NTP Servers
		//"0.europe.pool.ntp.org",
		//"1.europe.pool.ntp.org",
		//"2.europe.pool.ntp.org",
		//"3.europe.pool.ntp.org",

		// General NTP Servers
		//"pool.ntp.org",

		// https://www.chrony.eu/status
		"1.pool.chrony.eu",
		"2.pool.chrony.eu",
		"3.pool.chrony.eu",
		"4.pool.chrony.eu",

		// https://gist.github.com/mutin-sa/eea1c396b1e610a2da1e5550d94b0453

		// Google Public NTP [AS15169]:
		"time.google.com",
		"time1.google.com",
		"time2.google.com",
		"time3.google.com",
		"time4.google.com",
		// Cloudflare NTP [AS13335]:
		"time.cloudflare.com",
		// Facebook NTP [AS32934]:
		"time.facebook.com",
		"time1.facebook.com",
		"time2.facebook.com",
		"time3.facebook.com",
		"time4.facebook.com",
		"time5.facebook.com",
		// Microsoft NTP server [AS8075]:
		"time.windows.com",
		// Apple NTP server [AS714, AS6185]:
		"time.apple.com",
		"time1.apple.com",
		"time2.apple.com",
		"time3.apple.com",
		"time4.apple.com",
		"time5.apple.com",
		"time6.apple.com",
		"time7.apple.com",
		"time.euro.apple.com",
		// DEC/Compaq/HP:
		"clepsydra.dec.com", //"clepsydra.dec.com/clepsydra.labs.hp.com/clepsydra.hpl.hp.com/usno.labs.hp.com",
		// NIST Internet Time Service (ITS) [AS49, AS104]:
		"time-a-g.nist.gov",
		"time-b-g.nist.gov",
		"time-c-g.nist.gov",
		"time-d-g.nist.gov",
		"time-a-wwv.nist.gov",
		"time-b-wwv.nist.gov",
		"time-c-wwv.nist.gov",
		"time-d-wwv.nist.gov",
		"time-a-b.nist.gov",
		"time-b-b.nist.gov",
		"time-c-b.nist.gov",
		"time-d-b.nist.gov",
		"time.nist.gov",
		"time-e-b.nist.gov",
		"time-e-g.nist.gov",
		"time-e-wwv.nist.gov",
		"utcnist.colorado.edu",
		"utcnist2.colorado.edu",
		// VNIIFTRI:
		//# Stratum 1:
		"ntp1.vniiftri.ru",
		"ntp2.vniiftri.ru",
		"ntp3.vniiftri.ru",
		"ntp4.vniiftri.ru",
		"ntp.sstf.nsk.ru",
		"ntp1.niiftri.irkutsk.ru",
		"ntp2.niiftri.irkutsk.ru",
		"vniiftri.khv.ru",
		"vniiftri2.khv.ru",
		//# Stratum 2:
		"ntp21.vniiftri.ru",
		// Mobatime:
		//# Stratum 1:
		"ntp.mobatime.ru",
		// NTP SERVERS:
		"ntp0.ntp-servers.net",
		"ntp1.ntp-servers.net",
		"ntp2.ntp-servers.net",
		"ntp3.ntp-servers.net",
		"ntp4.ntp-servers.net",
		"ntp5.ntp-servers.net",
		"ntp6.ntp-servers.net",
		"ntp7.ntp-servers.net",
		//# Stratum 1:
		"ntp1.stratum1.ru",
		"ntp2.stratum1.ru",
		"ntp3.stratum1.ru",
		"ntp4.stratum1.ru",
		"ntp5.stratum1.ru",
		//# Stratum 2:
		"ntp1.stratum2.ru", //Москва
		"ntp2.stratum2.ru",
		"ntp3.stratum2.ru",
		"ntp4.stratum2.ru",
		"ntp5.stratum2.ru",
		// Stratum1:
		//# Stratum 1:
		"stratum1.net",
		// time.in.ua:
		//# Stratum 1:
		"ntp.time.in.ua",
		"ntp2.time.in.ua",
		//# Stratum 2:
		"ntp3.time.in.ua",
		// Company Delfa Co. Ltd. [AS8915]:
		"ntp.ru",
		// ACO.net [AS1853]:
		"ts1.aco.net",
		"ts2.aco.net",
		// Berkeley [AS25]:
		//# Stratum 1:
		"ntp1.net.berkeley.edu",
		"ntp2.net.berkeley.edu",
		// Georgia State University [AS10631]:
		"ntp.gsu.edu",
		// University of Saskatchewan [AS22950]:
		"tick.usask.ca",
		"tock.usask.ca",
		// NSU [AS3335]:
		//# Stratum 2:
		"ntp.nsu.ru",
		// ITAEC [AS41783]:
		"ntp.psn.ru",
		// RSU [AS47124]:
		//# Stratum 1:
		"ntp.rsu.edu.ru",
		// National Institute of Information and Communications Technology [AS9355]:
		"ntp.nict.jp",
		// NTT [AS2914]:
		"x.ns.gin.ntt.net",
		"y.ns.gin.ntt.net",
		// HE.net [AS6939]:
		"clock.nyc.he.net",
		"clock.sjc.he.net",
		// TRC Fiord [AS28917]:
		"ntp.fiord.ru",
		// Netnod NTP service [AS57021]:
		// Stratum 1:
		//Göteborg:
		"gbg1.ntp.se",
		"gbg2.ntp.se",
		//Malmö:
		"mmo1.ntp.se",
		"mmo2.ntp.se",
		//Stockholm:
		"sth1.ntp.se",
		"sth2.ntp.se",
		//Sundsvall:
		"svl1.ntp.se",
		"svl2.ntp.se",
		//Anycast address for nearest NTP server of the above:
		"ntp.se",
		// QiX NTP [AS14086]:
		"ntp.qix.ca",
		"ntp1.qix.ca",
		"ntp2.qix.ca",
		// YYCIX NTP [AS396515]:
		"ntp.yycix.ca",
		// MSK-IX NTP [AS43832]:
		//# Stratum 1:
		"ntp.ix.ru",
		// Hetzner Online [AS24940]:
		"ntp1.hetzner.de",
		"ntp2.hetzner.de",
		"ntp3.hetzner.de",
		// Trabia-Network [AS43289]:
		"time-a.as43289.net",
		"time-b.as43289.net",
		"time-c.as43289.net",
		// RIPE [AS3333]:
		"ntp.ripe.net",
		// Internet Systems Consortium [AS1280]:
		"clock.isc.org", //(prev ntp.isc.org)
		// TimeNL/SIDN Labs [AS1140]:
		"ntp.time.nl", // (ntp1.time.nl)
		// Kantonsschule Zug [AS34288]:
		"ntp0.as34288.net",
		"ntp1.as34288.net",
		// INTERNET MULTIFEED CO. [AS7521]:
		"ntp1.jst.mfeed.ad.jp",
		"ntp2.jst.mfeed.ad.jp",
		"ntp3.jst.mfeed.ad.jp",
		// Chinese Academy of Sciences Nation Time Service Center [AS4808, AS9808, AS23724]:
		"ntp.ntsc.ac.cn",
		// Nat Morris [AS30746]:
		//# Stratum 1:
		"ntp.nat.ms",
		// NTP Pool:
		"pool.ntp.org",
		"0.pool.ntp.org",
		"1.pool.ntp.org",
		"2.pool.ntp.org",
		"3.pool.ntp.org",
		"europe.pool.ntp.org",
		"0.europe.pool.ntp.org",
		"1.europe.pool.ntp.org",
		"2.europe.pool.ntp.org",
		"3.europe.pool.ntp.org",
		"asia.pool.ntp.org",
		"0.asia.pool.ntp.org",
		"1.asia.pool.ntp.org",
		"2.asia.pool.ntp.org",
		"3.asia.pool.ntp.org",
		"ru.pool.ntp.org",
		"0.ru.pool.ntp.org",
		"1.ru.pool.ntp.org",
		"2.ru.pool.ntp.org",
		"3.ru.pool.ntp.org",
		"north-america.pool.ntp.org",
		"0.north-america.pool.ntp.org",
		"1.north-america.pool.ntp.org",
		"2.north-america.pool.ntp.org",
		"3.north-america.pool.ntp.org",
		"0.gentoo.pool.ntp.org",
		"1.gentoo.pool.ntp.org",
		"2.gentoo.pool.ntp.org",
		"3.gentoo.pool.ntp.org",
		"0.arch.pool.ntp.org",
		"1.arch.pool.ntp.org",
		"2.arch.pool.ntp.org",
		"3.arch.pool.ntp.org",
		"0.fedora.pool.ntp.org",
		"1.fedora.pool.ntp.org",
		"2.fedora.pool.ntp.org",
		"3.fedora.pool.ntp.org",
		"0.opensuse.pool.ntp.org",
		"1.opensuse.pool.ntp.org",
		"2.opensuse.pool.ntp.org",
		"3.opensuse.pool.ntp.org",
		"0.centos.pool.ntp.org",
		"1.centos.pool.ntp.org",
		"2.centos.pool.ntp.org",
		"3.centos.pool.ntp.org",
		"0.debian.pool.ntp.org",
		"1.debian.pool.ntp.org",
		"2.debian.pool.ntp.org",
		"3.debian.pool.ntp.org",
		"0.askozia.pool.ntp.org",
		"1.askozia.pool.ntp.org",
		"2.askozia.pool.ntp.org",
		"3.askozia.pool.ntp.org",
		"0.freebsd.pool.ntp.org",
		"1.freebsd.pool.ntp.org",
		"2.freebsd.pool.ntp.org",
		"3.freebsd.pool.ntp.org",
		"0.netbsd.pool.ntp.org",
		"1.netbsd.pool.ntp.org",
		"2.netbsd.pool.ntp.org",
		"3.netbsd.pool.ntp.org",
		"0.openbsd.pool.ntp.org",
		"1.openbsd.pool.ntp.org",
		"2.openbsd.pool.ntp.org",
		"3.openbsd.pool.ntp.org",
		"0.dragonfly.pool.ntp.org",
		"1.dragonfly.pool.ntp.org",
		"2.dragonfly.pool.ntp.org",
		"3.dragonfly.pool.ntp.org",
		"0.pfsense.pool.ntp.org",
		"1.pfsense.pool.ntp.org",
		"2.pfsense.pool.ntp.org",
		"3.pfsense.pool.ntp.org",
		"0.opnsense.pool.ntp.org",
		"1.opnsense.pool.ntp.org",
		"2.opnsense.pool.ntp.org",
		"3.opnsense.pool.ntp.org",
		"0.smartos.pool.ntp.org",
		"1.smartos.pool.ntp.org",
		"2.smartos.pool.ntp.org",
		"3.smartos.pool.ntp.org",
		"0.android.pool.ntp.org",
		"1.android.pool.ntp.org",
		"2.android.pool.ntp.org",
		"3.android.pool.ntp.org",
		"0.amazon.pool.ntp.org",
		"1.amazon.pool.ntp.org",
		"2.amazon.pool.ntp.org",
		"3.amazon.pool.ntp.org",
		//# Other:
		// .mil:
		"tick.usno.navy.mil",
		"tock.usno.navy.mil",
		"ntp2.usno.navy.mil",
		// .edu:
		"utcnist.colorado.edu",
		"utcnist2.colorado.edu",
		"timekeeper.isi.edu",
		"rackety.udel.edu",
		"mizbeaver.udel.edu",
		"otc1.psu.edu",
		"gnomon.cc.columbia.edu",
		"navobs1.gatech.edu",
		"navobs1.wustl.edu",
		"now.okstate.edu",
		"ntp.colby.edu",
		"ntp-s1.cise.ufl.edu",
		"bonehed.lcs.mit.edu",
		"level1e.cs.unc.edu",
		"tick.ucla.edu",
		"tick.uh.edu",
		// .com:
		"ntpstm.netbone-digital.com",
		"nist1.symmetricom.com",
		"ntp.quintex.com",
		"ntp1.conectiv.com",
		"tock.usshc.com",
		// .net:
		"t2.timegps.net",
		"gps.layer42.net",
		"ntp-ca.stygium.net",
		"sesku.planeacion.net",
		"ntp0.nl.uu.net",
		"ntp1.nl.uu.net",
		"navobs1.oar.net",
		"ntp-galway.hea.net",
		// .org:
		"ntp1.ona.org",
		"ntp.your.org",
		"ntp.mrow.org",
		// .de:
		"time.fu-berlin.de",
		"ntps1-0.cs.tu-berlin.de",
		"ntps1-1.cs.tu-berlin.de",
		"ntps1-0.uni-erlangen.de",
		"ntps1-1.uni-erlangen.de",
		"ntp1.fau.de",
		"ntp2.fau.de",
		"ntp.dianacht.de",
		"zeit.fu-berlin.de",
		"ptbtime1.ptb.de",
		"ptbtime2.ptb.de",
		"rustime01.rus.uni-stuttgart.de",
		"rustime02.rus.uni-stuttgart.de",
		// .nl:
		"chime1.surfnet.nl",
		"ntp.vsl.nl",
		// .at:
		"asynchronos.iiss.at",
		// .cz:
		"ntp.nic.cz",
		"time.ufe.cz",
		// .pl:
		"ntp.fizyka.umk.pl",
		"tempus1.gum.gov.pl",
		"tempus2.gum.gov.pl",
		// .ro:
		"ntp1.usv.ro",
		"ntp3.usv.ro",
		// .se:
		"timehost.lysator.liu.se",
		"time1.stupi.se",
		// .ca:
		"time.nrc.ca",
		"clock.uregina.ca",
		// .mx:
		"cronos.cenam.mx",
		"ntp.lcf.mx",
		// .es:
		"hora.roa.es",
		"minuto.roa.es",
		// .it:
		"ntp1.inrim.it",
		"ntp2.inrim.it",
		// .be:
		"ntp1.oma.be",
		"ntp2.oma.be",
		// .hu:
		"ntp.atomki.mta.hu",
		// .eus:
		"ntp.i2t.ehu.eus",
		// .ch:
		"ntp.neel.ch",
		// .cn:
		"ntp.neu.edu.cn",
		// .jp:
		"ntp.nict.jp",
		// .br:
		"ntps1.pads.ufrj.br",
		// .cl:
		"ntp.shoa.cl",
		// .int:
		"time.esa.int",
		"time1.esa.int",

		/*
			http://support.ntp.org/bin/view/Servers/StratumOneTimeServers
			http://support.ntp.org/bin/view/Servers/StratumTwoTimeServers
			http://support.ntp.org/bin/view/Servers/NTPPoolServers
			http://www.pool.ntp.org/zone/@
			http://www.pool.ntp.org/zone/asia
			http://www.pool.ntp.org/zone/europe
			http://www.pool.ntp.org/zone/north-america
			http://www.pool.ntp.org/zone/oceania
			http://www.pool.ntp.org/zone/south-america
			https://time.nl/
			https://time.nl/index_en.html
			http://time.in.ua/
			https://www.chrony.eu/status
			https://www.ntp-server.de/ntp-server-deutschland/
			https://www.ntp-server.de/ntp-server-oesterreich/
			https://www.ntp-server.de/ntp-server-schweiz/
		*/

		// Add more NTP servers as needed
	}

	var serverInfoList []ServerInfo
	var accurateServers []ServerInfo

	for _, server := range ntpServers {
		response, err := ntp.Query(server)
		if err != nil {
			fmt.Printf("Error fetching time from %s: %v\n", server, err)
		} else {
			serverInfo := ServerInfo{
				Server:         server,
				RTT:            response.RTT,
				RootDispersion: response.RootDispersion,
			}
			if response.RootDispersion == 0 {
				accurateServers = append(accurateServers, serverInfo)
			}
			serverInfoList = append(serverInfoList, serverInfo)
			fmt.Printf("Server: %s, Speed (RTT): %v, Accuracy (RootDispersion): %v\n",
				server, response.RTT, response.RootDispersion)
		}
	}

	if len(accurateServers) > 0 {
		// Sort by RTT (fastest first)
		sort.Slice(accurateServers, func(i, j int) bool {
			return accurateServers[i].RTT < accurateServers[j].RTT
		})
	
		fmt.Println("\nNTP Servers with 100% Accuracy:")
		printTopN(accurateServers, len(accurateServers))
	} else {
		// Sort original list by RTT (fastest first)
		sort.Slice(serverInfoList, func(i, j int) bool {
			return serverInfoList[i].RTT < serverInfoList[j].RTT
		})
	
		fmt.Println("\nTop 20 Fastest NTP Servers:")
		printTopN(serverInfoList, 20)
	}

	// Sort by RTT (fastest first)
	sort.Slice(serverInfoList, func(i, j int) bool {
		return serverInfoList[i].RTT < serverInfoList[j].RTT
	})

	fmt.Println("Top 20 Fastest NTP Servers:")
	printTopN(serverInfoList, 20)
}

func printTopN(serverInfoList []ServerInfo, n int) {
	for i := 0; i < n && i < len(serverInfoList); i++ {
		fmt.Printf("%d. Server: %s, Speed (RTT): %v, Accuracy (RootDispersion): %v\n",
			i+1, serverInfoList[i].Server, serverInfoList[i].RTT, serverInfoList[i].RootDispersion)
	}
}
