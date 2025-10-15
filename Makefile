killall:
	killall leagueserver # ucenter snsserver
	echo "killall gameservering"
	@curl 127.0.0.1:19143/server/shutdown
	sleep 11
	echo "killall over"

deploy:
	#./snsserver >log/sns.log&
	#sleep 1
	@if pgrep -x "leagueserver" >/dev/null; then \
		echo "leagueserver already running"; \
	else \
		./leagueserver >log/lg.log & \
	fi

	sleep 1

	#./ucenter > log/uc.log&
	#sleep 1

	@if pgrep -x "gameserver" >/dev/null; then \
		echo "gameserver already running"; \
	else \
		./gameserver sandbox > log/gs.log & \
	fi

op:
	make killall
	#make rmpro
	make deploy

game:
	curl 127.0.0.1:19143/server/shutdown
	sleep 10
	echo "killall over"
	./gameserver sandbox > log/gs.log &

rmpro:
	cd csv && rm -rf product

rmlog:
	cd log && rm gameserver*.log leagueserver*.log ucenter*.log snsserver*.log webadmin_*.log balance*.log logmapobj.log