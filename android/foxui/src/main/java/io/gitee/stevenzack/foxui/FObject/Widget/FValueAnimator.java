package io.gitee.stevenzack.foxui.FObject.Widget;

import android.animation.ValueAnimator;

import org.json.JSONArray;
import org.json.JSONTokener;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;

public class FValueAnimator  extends FObject {
    public ValueAnimator v;

    public FValueAnimator(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new ValueAnimator();
    }

    @Override
    public String getAttr(String attr) {
        return "";
    }

    @Override
    public String setAttr(String attr, final String value, String value2) {
        switch (attr) {
            case "OfFloat":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    float[] vs = new float[array.length()];
                    for (int i=0;i<array.length();i++){
                        vs[i] = (float) array.getDouble(i);
                    }
                    v.setFloatValues(vs);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "OfInt":
                try {
                    JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
                    int[] vs = new int[array.length()];
                    for (int i=0;i<array.length();i++){
                        vs[i] =  array.getInt(i);
                    }
                    v.setIntValues(vs);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "Duration":
                try {
                    long d = Long.parseLong(value);
                    v.setDuration(d);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "ValueChangedListener":
                v.addUpdateListener(new ValueAnimator.AnimatorUpdateListener() {
                    @Override
                    public void onAnimationUpdate(ValueAnimator animation) {
//                        Faithdroid.triggerEventHandler(value, animation.getAnimatedValue().toString());
                        Fox.triggerFunction(parentController, value, animation.getAnimatedValue().toString(),"","");
                    }
                });
                break;
            case "Start":
                v.start();
                break;
        }
        return "";
    }
}
