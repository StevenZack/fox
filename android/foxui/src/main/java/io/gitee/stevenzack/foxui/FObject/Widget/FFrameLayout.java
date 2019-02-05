package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.v4.view.ViewCompat;
import android.view.View;
import android.widget.FrameLayout;

import org.json.JSONArray;
import org.json.JSONTokener;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FFrameLayout extends FObject {
    public FrameLayout v;
    public FFrameLayout(FoxActivity controller) {
        parentController = controller;
        v = new FrameLayout(parentController);
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
            // --------------------------------------------------
        }
        return "";
    }

    @Override
    public String setAttr(String attr, final String value,final String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str!=null) {
            return str;
        }
        switch (attr){
            // ----------------------------------------------------------------------------
            case "AddView":
                final String childVid = value;
                FObject f = parentController.viewmap.get(childVid);
                v.addView(f.view,parseLP(f));
                break;
            case "AddViewAt":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    int pos = Integer.parseInt(array.getString(0));
                    String vid = array.getString(1);
                    FObject f1 = parentController.viewmap.get(vid);
                    v.addView(f1.view,pos,parseLP(f1));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "OnClick":
                v.setOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View view) {
                        Fox.triggerFunction(parentController,value, "","","");
                    }
                });
                break;
        }
        return "";
    }


    public static FrameLayout.LayoutParams parseLP(FObject f) {
        FrameLayout.LayoutParams lp = new FrameLayout.LayoutParams(f.size[0], f.size[1]);
        lp.gravity=f.layoutGravity;
        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        return lp;
    }
}
