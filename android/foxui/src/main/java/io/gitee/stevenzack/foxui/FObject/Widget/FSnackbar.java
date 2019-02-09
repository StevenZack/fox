package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.design.widget.Snackbar;
import android.view.View;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FSnackbar extends FObject {
    public Snackbar v;

    public FSnackbar(FoxActivity foxActivity, View anchor) {
        parentController = foxActivity;
        v = Snackbar.make(anchor, "", Snackbar.LENGTH_SHORT);
    }

    @Override
    public String getAttr(String attr) {
        return "";
    }

    @Override
    public String setAttr(String attr, String value, final String value2) {
        switch (attr) {
            case "Text":
                v.setText(value);
                break;
            case "Action":
                v.setAction(value, new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Fox.triggerFunction(parentController, value2, "", "", "");
                    }
                });
                break;
            case "Show":
                v.show();
                break;
        }
        return "";
    }
}
