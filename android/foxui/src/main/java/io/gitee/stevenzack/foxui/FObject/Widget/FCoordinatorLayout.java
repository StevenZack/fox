package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.design.widget.CoordinatorLayout;
import android.support.v4.view.ViewCompat;
import android.util.Log;
import android.view.ViewGroup;
import android.widget.FrameLayout;

import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FCoordinatorLayout extends FObject  {
    public CoordinatorLayout v;

    public FCoordinatorLayout(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new CoordinatorLayout(foxActivity);
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
    public String setAttr(String attr, String value, String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "AddView":
                FObject fObject = parentController.viewmap.get(value);
                if (fObject != null) {
                    v.addView(fObject.view,parseLP(fObject));
                }
                break;
            case "AddViewAt":
                FObject fObject1 = parentController.viewmap.get(value);
                if (fObject1 != null) {
                    v.addView(fObject1.view,Integer.parseInt(value2),parseLP(fObject1));
                }else{
                    Log.d(TAG, "setAttr: fobject is null");
                }
                break;
            case "Append":
                String[] vids = value.split(",");
                for (int i = 0; i < vids.length; i++) {
                    FObject object=parentController.viewmap.get(vids[i]);
                    if (object != null) {
                        v.addView(object.view,parseLP(object));
                    }else{
                        Log.d(TAG, "setAttr: object is null");
                    }
                }
                break;
        }
        return "";
    }

    private ViewGroup.LayoutParams parseLP(FObject f) {
        CoordinatorLayout.LayoutParams lp = new CoordinatorLayout.LayoutParams(f.size[0], f.size[1]);
        lp.gravity=f.layoutGravity;
        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        return lp;
    }
}
