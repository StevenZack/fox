package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.v4.view.ViewCompat;
import android.support.v7.widget.LinearLayoutCompat;
import android.util.Log;
import android.widget.LinearLayout;

import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FBox extends FObject {
    public LinearLayoutCompat v;

    public FBox(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new LinearLayoutCompat(parentController);
        view = v;
        v.setId(ViewCompat.generateViewId());
        v.setOrientation(LinearLayoutCompat.VERTICAL);
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
            case "Orientation":
                if (value.equals("HORIZONTAL")) {
                    v.setOrientation(LinearLayout.HORIZONTAL);
                }else {
                    v.setOrientation(LinearLayout.VERTICAL);
                }
                break;
        }
        return "";
    }

    public static LinearLayoutCompat.LayoutParams parseLP(FObject f) {
        LinearLayoutCompat.LayoutParams lp = new LinearLayoutCompat.LayoutParams(f.size[0], f.size[1]);
        lp.gravity=f.layoutGravity;
        lp.weight=f.layoutWeight;
        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        return lp;
    }
}
