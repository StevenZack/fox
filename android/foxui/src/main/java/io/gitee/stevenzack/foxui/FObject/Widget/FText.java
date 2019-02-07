package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

import static io.gitee.stevenzack.foxui.Toolkit.dp2pixel;

public class FText extends FObject {
    public TextView v;
    public FText(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new TextView(parentController);
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
            case "Text":
                return v.getText().toString();
        }
        return "";
    }

    @Override
    public String setAttr(String attr, final String value, final String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "Text":
                v.setText(value);
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Fox.triggerFunction(null, value, "", "","");
                    }
                });
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
        }
        return "";
    }
}
