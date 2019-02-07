package io.gitee.stevenzack.foxui.FObject;

import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.constraint.ConstraintLayout;
import android.support.v4.view.ViewCompat;
import android.util.Log;
import android.view.View;
import android.view.ViewGroup;

import org.json.JSONArray;
import org.json.JSONTokener;

import java.util.HashMap;
import java.util.Map;

import io.gitee.stevenzack.foxui.FObject.Widget.FConstraintLayout;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.Toolkit;

import static io.gitee.stevenzack.foxui.Toolkit.dp2pixel;

public abstract class FObject {
    public String vid ,vtype ;
    public View view;
    public FoxActivity parentController;
    public int[] size=new int[]{-2,-2};
    public int[] margin=new int[4];
    public int layoutGravity;
    public float layoutWeight;
    protected String TAG="FObject";
    public Map<String, ConstraintInterface> afterConstraintFuncs = new HashMap<>();
    public interface ConstraintInterface{
        void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp);
    }

    public String getAttr(String attr){
        return null;
    }
    public String setAttr(String attr, final String value, String value2){
        switch (attr) {
            case "BackgroundColor":
                setBackgroundColor(value);
                break;
            case "Background":
                setBackground(value);
                break;
            case "Foreground":
                setForeground(value);
                break;
            case "Size":
                parseSize(value,value2);
                break;
            case "X":
                setX(value);
                break;
            case "Y":
                setY(value);
                break;
            case "PivotX":
                setPivotX(value);
                break;
            case "PivotY":
                setPivotY(value);
                break;
            case "ScaleX":
                setScaleX(value);
                break;
            case "ScaleY":
                setScaleY(value);
                break;
            case "Rotation":
                setRotation(value);
                break;
            case "Visibility":
                setVisibility(value);
                break;
            case "Padding":
                setPadding(value);
                break;
            case "Margin":
                setMargin(value);
                break;
            case "LayoutGravity":
                setLayoutGravity(value);
                break;
            case "Elevation":
                setElevation(value);
                break;
            case "LayoutWeight":
                setLayoutWeight(value);
                break;
            case "Clickable":
                view.setClickable(value.equals("true"));
                break;
            // constraint
            case "Top2TopOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.topToTop = parent.view.getId();
                            return;
                        }
                        lp.topToTop = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "Top2BottomOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.bottomToBottom = parent.view.getId();
                            return;
                        }
                        lp.topToBottom = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "Bottom2BottomOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.bottomToBottom = parent.view.getId();
                            return;
                        }
                        lp.bottomToBottom = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "Bottom2TopOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.bottomToTop = parent.view.getId();
                            return;
                        }
                        lp.bottomToTop = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "Left2LeftOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.leftToLeft = parent.view.getId();
                            return;
                        }
                        lp.leftToLeft = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "Right2RightOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.rightToRight = parent.view.getId();
                            return;
                        }
                        lp.rightToRight = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "Left2RightOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.leftToRight = parent.view.getId();
                            return;
                        }
                        lp.leftToRight = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "Right2LeftOf":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("_Parent_")) {
                            lp.rightToLeft = parent.view.getId();
                            return;
                        }
                        lp.rightToLeft = parentController.viewmap.get(value).view.getId();
                    }
                });
                break;
            case "CenterX":
                setAttr("Left2LeftOf", "_Parent_","");
                setAttr("Right2RightOf", "_Parent_","");
                break;
            case "CenterY":
                setAttr("Top2TopOf", "_Parent_","");
                setAttr("Bottom2BottomOf", "_Parent_","");
                break;
            case "WidthPercent":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("1")) {
                            lp.width = -1;//matchParent
                            return;
                        }
                        lp.width = 0;
                        lp.matchConstraintPercentWidth = Float.parseFloat(value);
                    }
                });
                break;
            case "HeightPercent":
                afterConstraintFuncs.put(attr, new ConstraintInterface() {
                    @Override
                    public void addConstraint(FConstraintLayout parent, ConstraintLayout.LayoutParams lp) {
                        if (value.equals("1")) {
                            lp.height = -1;//matchParent
                            return;
                        }
                        lp.height = 0;
                        lp.matchConstraintPercentHeight = Float.parseFloat(value);
                    }
                });
                break;
            default:
                return null;
        }
        return "";
    }
    void setBackgroundColor(String value) {
        if (value==null)
            return;
        if (value.equals("#0000000")) {
            view.setBackgroundColor(Color.TRANSPARENT);
            return;
        }
        try {
            view.setBackgroundColor(Color.parseColor(value));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setBackground(String value) {
        if (value==null)
            return;
        if (value.startsWith("#")) {
            setBackgroundColor(value);
            return;
        }
        Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
            @Override
            public void onDrawableReady(Drawable draw) {
                if (draw == null) {
                    return;
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                    view.setBackground(draw);
                }else {
                    view.setBackgroundDrawable(draw);
                }
            }
        });
        if (value.equals("RippleEffect")) {
            view.setClickable(true);
        }
    }

    void setForeground(String value) {
        if (value == null) {
            return;
        }
        Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
            @Override
            public void onDrawableReady(Drawable draw) {
                if (draw == null) {
                    return;
                }
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    view.setForeground(draw);
                }
            }
        });

        if (value.equals("RippleEffect")) {
            view.setClickable(true);
        }
    }
    protected void parseSize(String value1,String value2) {
        long width=Integer.parseInt(value1), height=Integer.parseInt(value2);
        ViewGroup.LayoutParams p = view.getLayoutParams();
        if (p == null) {
            p=new ViewGroup.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT);
        }
        if (width == -1) {
            p.width= ViewGroup.LayoutParams.WRAP_CONTENT;
        } else if (width == -2) {
            p.width= ViewGroup.LayoutParams.MATCH_PARENT;
        }else{
            p.width = (int) dp2pixel(parentController, width);
        }
        if (height == -1) {
            p.height = ViewGroup.LayoutParams.WRAP_CONTENT;
        } else if (height == -2) {
            p.height = ViewGroup.LayoutParams.MATCH_PARENT;
        }else{
            p.height = (int) dp2pixel(parentController, height);
        }
        size[0]=p.width;
        size[1]=p.height;
    }
    void setX(String value) {
        try {
            float f = dp2pixel(parentController, Float.parseFloat(value));
            view.setX(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setY(String value) {
        try {
            float f = dp2pixel(parentController, Float.parseFloat(value));
            view.setY(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setPivotX(String value) {
        try {
            float f = dp2pixel(parentController, Float.parseFloat(value));
            view.setPivotX(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setPivotY(String value) {
        try {
            float f = dp2pixel(parentController, Float.parseFloat(value));
            view.setPivotY(f);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setScaleX(String value) {
        try {
//            float f = dp2pixel(parentController, Float.parseFloat(value));
            view.setScaleX(Float.parseFloat(value));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setScaleY(String value) {
        try {
//            float f = dp2pixel(parentController, Float.parseFloat(value));

            view.setScaleY(Float.parseFloat(value));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setRotation(String value) {
        try {
            view.setRotation(Float.parseFloat(value));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setVisibility(String value) {
        int vsb=View.VISIBLE;
        if (value.equals("INVISIBLE")){
            vsb=View.INVISIBLE;
        } else if (value.equals("GONE")) {
            vsb=View.GONE;
        }
        view.setVisibility(vsb);
    }
    String getVisibility(){
        int vsb=view.getVisibility();
        if (vsb== View.VISIBLE) {
            return "VISIBLE";
        } else if (vsb == View.GONE) {
            return "GONE";
        }
        return "INVISIBLE";
    }
    void setPadding(String value) {
        try {
            JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
            int left= (int) dp2pixel(parentController,array.getLong(0));
            int top= (int) dp2pixel(parentController,array.getLong(1));
            int right= (int) dp2pixel(parentController,array.getLong(2));
            int bottom= (int) dp2pixel(parentController,array.getLong(3));
            view.setPadding(left,top,right,bottom);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setMargin(String value) {
        try {
            JSONArray array = (JSONArray) (new JSONTokener(value).nextValue());
            margin[0] =(int)dp2pixel(parentController, array.getInt(0));
            margin[1] =(int)dp2pixel(parentController, array.getInt(1));
            margin[2] =(int)dp2pixel(parentController, array.getInt(2));
            margin[3] =(int)dp2pixel(parentController, array.getInt(3));
//            if (view.getLayoutParams() instanceof LinearLayout.LayoutParams) {
//                LinearLayout.MarginLayoutParams params = new LinearLayout.MarginLayoutParams(view.getLayoutParams());
//                params.setMargins((int) (dp2pixel(parentController, array.getInt(0))),
//                        (int) (dp2pixel(parentController, array.getInt(2))),
//                        (int) (dp2pixel(parentController, array.getInt(1))),
//                        (int) (dp2pixel(parentController, array.getInt(3))));
//                view.setLayoutParams(params);
//            } else if (view.getLayoutParams() instanceof FrameLayout.LayoutParams) {
//                FrameLayout.MarginLayoutParams params = new FrameLayout.MarginLayoutParams(view.getLayoutParams());
//                params.setMargins((int) (dp2pixel(parentController, array.getInt(0))),
//                        (int) (dp2pixel(parentController, array.getInt(2))),
//                        (int) (dp2pixel(parentController, array.getInt(1))),
//                        (int) (dp2pixel(parentController, array.getInt(3))));
//                view.setLayoutParams(params);
//            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setLayoutGravity(String value) {
        try {
            layoutGravity = Integer.parseInt(value);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    void setElevation(String value) {
        try {
            float f = Float.parseFloat(value);
            ViewCompat.setElevation(view,f);
            Log.d(TAG, "setElevation: "+value);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    void setLayoutWeight(String value) {
        try {
            layoutWeight = Float.parseFloat(value);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}
