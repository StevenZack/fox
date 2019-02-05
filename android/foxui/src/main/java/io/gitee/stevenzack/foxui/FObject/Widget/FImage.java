package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.drawable.Drawable;
import android.support.v4.view.ViewCompat;
import android.support.v7.widget.AppCompatImageView;
import android.util.Log;
import android.view.View;
import android.widget.Button;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.Toolkit;

public class FImage extends FObject {
    public AppCompatImageView v;
    public FImage(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new AppCompatImageView(parentController);
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
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Fox.triggerFunction(null, value, "", "","");
                    }
                });
                break;
            case "Src":
                Log.d(TAG, "setAttr: Src="+value);
                Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
                    @Override
                    public void onDrawableReady(Drawable d) {
                        v.setImageDrawable(d);
                    }
                });
                break;
        }
        return "";
    }
}
