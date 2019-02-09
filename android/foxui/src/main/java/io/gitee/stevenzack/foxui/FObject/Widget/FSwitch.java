package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.v4.view.ViewCompat;
import android.support.v7.widget.SwitchCompat;
import android.widget.CompoundButton;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FSwitch extends FObject {
    public SwitchCompat v;

    public FSwitch(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new SwitchCompat(parentController);
        v.setId(ViewCompat.generateViewId());
        view=v;
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "Checked":
                return String.valueOf(v.isChecked());
        }
        return "";
    }

    @Override
    public String setAttr(String attr, final String value, String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "OnChange":
                v.setOnCheckedChangeListener(new CompoundButton.OnCheckedChangeListener() {
                    @Override
                    public void onCheckedChanged(CompoundButton buttonView, boolean isChecked) {
                        Fox.triggerFunction(parentController, value, "", "", "");
                    }
                });
                break;
        }
        return "";
    }
}
