package io.gitee.stevenzack.foxui.FObject.Widget;

import android.app.AlertDialog;
import android.content.DialogInterface;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FDialog extends FObject {
    public AlertDialog.Builder v;
    public DialogInterface di;

    public FDialog(FoxActivity activity) {
        parentController = activity;
        v = new AlertDialog.Builder(parentController);
    }

    @Override
    public String getAttr(String attr) {
        return "";
    }

    @Override
    public String setAttr(String attr, String value, final String value2) {
        if (value == null)
            return "";
        switch (attr) {
            case "Title":
                v.setTitle(value);
                break;
            case "View":
                if (parentController.viewmap.containsKey(value)) {
                    v.setView(parentController.viewmap.get(value).view);
                }
                break;
            case "PositiveButton":
                try {
                    v.setPositiveButton(value, new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialog, int which) {
                            Fox.triggerFunction(parentController, value2, "", "", "");
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "NegativeButton":
                try {
                    v.setNegativeButton(value, new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface dialog, int which) {
                            Fox.triggerFunction(parentController, value2, "", "", "");
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Show":
                di = v.show();
                break;
            case "Dismiss":
                if (di != null) {
                    di.dismiss();
                }
                break;
        }
        return "";
    }
}
