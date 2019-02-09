package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.widget.CheckBox;
import android.widget.CompoundButton;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

import static io.gitee.stevenzack.foxui.Toolkit.dp2pixel;

public class FCheck extends FObject {
    public CheckBox v;
    private String onChange;

    public FCheck(FoxActivity activity) {
        parentController = activity;
        v = new CheckBox(parentController);
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
            // -------------------------------------------------------------------
            case "Text":
                v.setText(value);
                break;
            case "TextColor":
                try {
                    v.setTextColor(Color.parseColor(value));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "TextSize":
                try {
                    v.setTextSize(dp2pixel(parentController,Float.valueOf(value)));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Enabled":
                if (value.equals("true")) {
                    v.setEnabled(true);
                } else {
                    v.setEnabled(false);
                }
                break;
            case "OnChange":
                if (onChange==null ){
                    v.setOnCheckedChangeListener(new CompoundButton.OnCheckedChangeListener() {
                        @Override
                        public void onCheckedChanged(CompoundButton buttonView, boolean isChecked) {
                            Fox.triggerFunction(parentController,onChange, String.valueOf(isChecked),"","");
                        }
                    });
                }
                onChange=value;
                break;
            case "Checked":
                v.setChecked(value.equals("true"));
                break;
        }
        return "";
    }
}
