package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.drawable.Drawable;
import android.support.v4.view.ViewCompat;
import android.support.v7.widget.AppCompatImageView;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.ImageView;

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
                Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
                    @Override
                    public void onDrawableReady(Drawable d) {
                        v.setImageDrawable(d);
                    }
                });
                break;
            case "ScaleType":
                v.setScaleType(parseScaleType(value));
                break;
        }
        return "";
    }
    private ImageView.ScaleType parseScaleType(String value) {
        switch (value) {
            case "CenterCrop":
                return ImageView.ScaleType.CENTER_CROP;
            case "CenterInside":
                return ImageView.ScaleType.CENTER_INSIDE;
            case "FitCenter":
                return ImageView.ScaleType.FIT_CENTER;
            case "FitStart":
                return ImageView.ScaleType.FIT_START;
            case "FitEnd":
                return ImageView.ScaleType.FIT_END;
            case "FitXY":
                return ImageView.ScaleType.FIT_XY;
            case "Matrix":
                return ImageView.ScaleType.MATRIX;
            default:
                return ImageView.ScaleType.CENTER;
        }
    }
}
