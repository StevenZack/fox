package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.constraint.ConstraintLayout;
import android.support.v4.view.ViewCompat;

import java.util.Map;

import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FConstraintLayout extends FObject {
    public ConstraintLayout v;

    public FConstraintLayout(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new ConstraintLayout(parentController);
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
            case "Append":
                String[] vids = value.split(",");
                for (int i=0;i<vids.length;i++) {
                    FObject f= parentController.viewmap.get(vids[i]);
                    ConstraintLayout.LayoutParams lp = new ConstraintLayout.LayoutParams(f.size[0], f.size[1]);
                    for (Map.Entry<String, ConstraintInterface> entry : f.afterConstraintFuncs.entrySet()) {
                        entry.getValue().addConstraint(this,lp);
                    }
                    lp.leftMargin = f.margin[0];
                    lp.topMargin = f.margin[1];
                    lp.rightMargin = f.margin[2];
                    lp.bottomMargin = f.margin[3];
                    v.addView(parentController.viewmap.get(vids[i]).view,lp);
                }
                break;
            case "AddView":
                addView(value, value2);
                break;
            case "AddViewAt":
                addView(value,value2);
        }
        return "";
    }

    private void addView(String value, String value2) {
        FObject f= parentController.viewmap.get(value);
        ConstraintLayout.LayoutParams lp = new ConstraintLayout.LayoutParams(f.size[0], f.size[1]);
        for (Map.Entry<String, ConstraintInterface> entry : f.afterConstraintFuncs.entrySet()) {
            entry.getValue().addConstraint(this,lp);
        }
        lp.leftMargin = f.margin[0];
        lp.topMargin = f.margin[1];
        lp.rightMargin = f.margin[2];
        lp.bottomMargin = f.margin[3];
        v.addView(parentController.viewmap.get(value).view,lp);
    }
}
