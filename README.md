# settime-restart
 When there are a lot of programs running on a server, settime-restart can help you to restart them regularly.
# only support windows 
1.
config.json�����ļ�:
{
	"program":[
		{
			"name":"��վӦ��", //������������
			"startcmd":"D:\\soft\\tomcat\\bin\\startup.bat", //��������
			"findname":"java_web", //��������
			"cron":"20 14 13 ? * TUE" //����ʱ�䣬��С���Ϊ1min����		
		}
	]	
}

2.
��ʱ����tomcat
(1).�޸� startup.bat 
CATALINA_HOMEΪ��ǰtomcat·�� ��:set CATALINA_HOME=D:\soft\tomcat2

(2).�޸� setclasspath.bat 
copy %JRE_HOME%\bin\java.exe %JRE_HOME%\bin\java_web.exe
set _RUNJAVA="%JRE_HOME%\bin\java_web.exe"