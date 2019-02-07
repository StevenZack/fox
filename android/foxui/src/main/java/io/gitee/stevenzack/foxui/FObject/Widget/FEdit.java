package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.Color;
import android.support.v4.view.ViewCompat;
import android.text.InputFilter;
import android.text.InputType;
import android.widget.EditText;

import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

import static io.gitee.stevenzack.foxui.Toolkit.dp2pixel;

public class FEdit extends FObject {
    public EditText v;

    public FEdit(FoxActivity activity) {
        parentController = activity;
        v = new EditText(activity);
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
            case "Text":
                return v.getText().toString();
        }
        return "";
    }

    @Override
    public String setAttr(String attr, final String value,final String value2) {
        String str=super.setAttr(attr, value,value2);
        if (str!=null) {
            return str;
        }
        switch (attr){
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
                    v.setTextSize(dp2pixel(parentController, Float.valueOf(value)));
                } catch (Exception e) {
                    e.printStackTrace();
                    return "";
                }
                break;
            case "InputType":
                int it;
                if (value.equals("Text")) {
                    it= InputType.TYPE_CLASS_TEXT;
                } else if (value.equals("Number")) {
                    it = InputType.TYPE_CLASS_NUMBER;
                } else if (value.equals("Password")) {
                    it=InputType.TYPE_CLASS_TEXT|InputType.TYPE_TEXT_VARIATION_WEB_PASSWORD;
                } else  {
                    it = InputType.TYPE_TEXT_VARIATION_VISIBLE_PASSWORD;
                }
                v.setInputType(it);
                break;
            case "MaxLines":
                v.setMaxLines(Integer.parseInt(value));
                break;
            case "MaxEms":
                v.setMaxEms(Integer.parseInt(value));
                break;
            case "Hint":
                v.setHint(value);
                break;
            case "MaxLength":
                v.setFilters(new InputFilter[] {new InputFilter.LengthFilter(Integer.parseInt(value))});
                break;
        }
        return "";
    }
}
