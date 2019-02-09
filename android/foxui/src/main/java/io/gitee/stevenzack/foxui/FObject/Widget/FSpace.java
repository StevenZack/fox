package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.v4.view.ViewCompat;
import android.widget.Space;

import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FSpace extends FObject {
    public Space v;

    public FSpace(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new Space(foxActivity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        layoutWeight=1;
    }

    @Override
    public String getAttr(String attr) {
        return super.getAttr(attr);
    }

    @Override
    public String setAttr(String attr, String value, String value2) {
        return super.setAttr(attr, value, value2);
    }
}
