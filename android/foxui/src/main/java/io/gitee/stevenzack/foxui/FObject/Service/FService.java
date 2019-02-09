package io.gitee.stevenzack.foxui.FObject.Service;

import android.content.Context;
import android.content.Intent;

import io.gitee.stevenzack.foxui.CoreService;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.R;

public class FService extends FObject {
    public FService(FoxActivity foxActivity) {
        parentController = foxActivity;
    }

    public static String getAttrStaticly(String attr) {
        return "";
    }

    public static void setAttr(Context context, String attr, final String value) {
        switch (attr) {
            case "OnCreate":
                CoreService.onCreate = value;
                break;
            case "QuitButton":
                CoreService.quitButton = value;
                break;
            case "OnQuit":
                CoreService.onQuitClick =value;
                break;
            case "Title":
                CoreService.ntf_title=value;
                if (CoreService.notification!=null){
                    CoreService.getInstance().handler.post(new Runnable() {
                        @Override
                        public void run() {
                            CoreService.notification.contentView.setTextViewText(R.id.ntf_title,value);
                            CoreService.getInstance().startForeground(1,CoreService.notification);
                        }
                    });
                }
                break;
            case "SubTitle":
                CoreService.ntf_subtitle = value;
                if (CoreService.notification != null) {
                    CoreService.getInstance().handler.post(new Runnable() {
                        @Override
                        public void run() {
                            CoreService.notification.contentView.setTextViewText(R.id.ntf_subtitle,value);
                            CoreService.getInstance().startForeground(1,CoreService.notification);
                        }
                    });
                }
                break;
            case "Show":
                Intent intent = new Intent(context, CoreService.class);
                context.startService(intent);
                break;
            case "FinishAllActivity":
                CoreService.getInstance().finishAllActivity();
                break;
            case "KillAll":
                CoreService.killAll();
                break;
            case "Stop":
                CoreService.getInstance().stopService();
                break;
        }
    }
}
