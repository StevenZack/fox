package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.widget.RadioButton;

import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

import static io.gitee.stevenzack.foxui.Toolkit.dp2pixel;

public class FRadio extends FObject {
    public RadioButton v;

    public FRadio(FoxActivity activity) {
        parentController = activity;
        v = new RadioButton(activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // ------------------------------------------
            case "IsSelected":
                return String.valueOf(v.isSelected());
            case "Text":
                return v.getText().toString();
        }
        return "";
    }

    @Override
    public String setAttr(String attr, String value, String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "Text":
                v.setText(value);
                break;
            case "Selected":
                v.setSelected(value.equals("true"));
                break;
        }
        return "";
    }
}
