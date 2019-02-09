package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.util.Log;
import android.view.View;
import android.widget.Button;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FButton extends FObject {
    public Button v;
    public FButton(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new Button(parentController);
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
            case "TextColor":
                v.setTextColor(Color.parseColor(value));
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Log.d(TAG, "onClick: clicked");
                        Fox.triggerFunction(parentController, value, "", "","");
                    }
                });
                break;
        }
        return "";
    }
}
