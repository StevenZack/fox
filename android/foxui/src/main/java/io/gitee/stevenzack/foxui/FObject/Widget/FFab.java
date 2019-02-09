package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.drawable.Drawable;
import android.support.design.widget.FloatingActionButton;
import android.support.v4.view.ViewCompat;
import android.view.View;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.Toolkit;

public class FFab extends FObject {
    public FloatingActionButton v;

    public FFab(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new FloatingActionButton(foxActivity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        setElevation("8");
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr) {

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
            // -------------------------------------------------------------------
            case "Icon":
                Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
                    @Override
                    public void onDrawableReady(Drawable drawable) {
                        v.setImageDrawable(drawable);
                    }
                });
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Fox.triggerFunction(parentController, value, "", "", "");
                    }
                });
                break;
        }
        return "";
    }
}
