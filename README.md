# settime-restart
 When there are a lot of programs running on a server, settime-restart can help you to restart them regularly.
# only support windows 
1.
config.json配置文件:
<pre>
{
	"program":[
		{
			"name":"网站应用", //启动任务名称
			"startcmd":"D:\\soft\\tomcat\\bin\\startup.bat", //启动命令
			"findname":"java_web", //进程名称
			"cron":"20 14 13 ? * TUE" //重启时间，最小间隔为1min以上		
		}
	]	
}
</pre>
2.
<pre>
定时重启tomcat
(1).修改 startup.bat 
CATALINA_HOME为当前tomcat路径 如:set CATALINA_HOME=D:\soft\tomcat2

(2).修改 setclasspath.bat 
copy %JRE_HOME%\bin\java.exe %JRE_HOME%\bin\java_web.exe
set _RUNJAVA="%JRE_HOME%\bin\java_web.exe"
</pre>
