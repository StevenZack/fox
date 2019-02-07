package io.gitee.stevenzack.foxui.FObject.Widget;

import android.support.v4.view.ViewCompat;
import android.widget.RadioGroup;

import java.util.HashMap;
import java.util.Map;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FRadioGroup extends FObject  {
    public RadioGroup v;
    public Map<Integer, String> idMap = new HashMap<>();

    public FRadioGroup(FoxActivity activity) {
        parentController = activity;
        v = new RadioGroup(activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        v.setOrientation(RadioGroup.VERTICAL);
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            // --------------------------------------------------
            case "Orientation":
                if (v.getOrientation()==RadioGroup.HORIZONTAL)
                    return "HORIZONTAL";
                else
                    return "VERTICAL";
            case "CheckedRadioButtonId":
                return idMap.get(v.getCheckedRadioButtonId());
        }
        return null;
    }

    @Override
    public String setAttr(String attr, final String value, String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            // ----------------------------------------------------------------------------
            case "Orientation":
                if (value.equals("HORIZONTAL")) {
                    v.setOrientation(RadioGroup.HORIZONTAL);
                } else {
                    v.setOrientation(RadioGroup.VERTICAL);
                }
                break;
            case "AddView":
                final String childVid = value;
                FObject f = parentController.viewmap.get(childVid);
                RadioGroup.LayoutParams lp = new RadioGroup.LayoutParams(f.size[0], f.size[1]);
                lp.gravity = f.layoutGravity;
                lp.weight = f.layoutWeight;
                lp.leftMargin = f.margin[0];
                lp.topMargin = f.margin[1];
                lp.rightMargin = f.margin[2];
                lp.bottomMargin = f.margin[3];
                v.addView(f.view,lp);
                idMap.put(f.view.getId(),f.vid);
                break;
            case "OnCheckedChange":
                v.setOnCheckedChangeListener(new RadioGroup.OnCheckedChangeListener() {
                    @Override
                    public void onCheckedChanged(RadioGroup group, int checkedId) {
                        String checkedVid = idMap.get(checkedId);
                        Fox.triggerFunction(parentController, value, checkedVid,"","");
                    }
                });
                break;
        }
        return "";
    }
}
