package io.gitee.stevenzack.foxui;

import android.content.Intent;
import android.graphics.drawable.Drawable;
import android.support.v4.graphics.drawable.DrawableCompat;
import android.support.v4.widget.DrawerLayout;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.util.Log;

import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.AbstractSequentialList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;

import fox.Fox;
import fox.IActivity;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FObject.Widget.FBox;
import io.gitee.stevenzack.foxui.FObject.Widget.FButton;
import io.gitee.stevenzack.foxui.FObject.Widget.FConstraintLayout;
import io.gitee.stevenzack.foxui.FObject.Widget.FDialog;
import io.gitee.stevenzack.foxui.FObject.Widget.FFrameLayout;
import io.gitee.stevenzack.foxui.FObject.Widget.FImage;
import io.gitee.stevenzack.foxui.FObject.Widget.FListView;
import io.gitee.stevenzack.foxui.FObject.Widget.FText;

public class FoxActivity extends AppCompatActivity implements IActivity {
    public Map<String, Drawable> drawableMap=new HashMap<>();
    private String activityId="MainActivity";
    public DrawerLayout rootCtn;
    public Map<String, FObject> viewmap = new HashMap<>();
    private String TAG="FoxActivity";

    protected void mainFoxUI(DrawerLayout rootCtn) {//only called by MainActivity (the entry port of Go Code
        this.rootCtn=rootCtn;
        Fox.main(this);
    }
    protected void startFoxUI(DrawerLayout rootCtn) {// called by those second activity
        this.rootCtn=rootCtn;
        try {
            String intentJson = Toolkit.handleIntent(this, getIntent());
            activityId=getIntent().getStringExtra("FActivityId");
            //oncreate
            Fox.triggerFunction(this, "StartActivity", activityId, "OnCreate", intentJson);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Override
    public String getActivityId() {
        return activityId;
    }

    @Override
    public String getAttr(String vid, String attr) {
        if (vid.equals("Activity")) {
            return getActivityAttr(attr);
        }
        FObject fObject = viewmap.get(vid);
        if (fObject != null) {
            return fObject.getAttr(attr);
        }
        return "";
    }


    @Override
    public String newObject(String vtype, String vid) {
        if (vtype.equals("Activity")) {
            return newActivity(vid);
        }
        FObject fObject=null;
        Log.d(TAG, "newObject: "+vtype);
        switch (vtype) {
            case "Box":
                fObject = new FBox(this);
                break;
            case "Button":
                fObject=new FButton(this);
                break;
            case "Text":
                fObject = new FText(this);
                break;
            case "Image":
                fObject = new FImage(this);
                break;
            case "ConstraintBox":
                fObject = new FConstraintLayout(this);
                break;
            case "FrameBox":
                fObject = new FFrameLayout(this);
                break;
            case "VList":
                LinearLayoutManager lmv = new LinearLayoutManager(this);
                lmv.setOrientation(LinearLayoutManager.VERTICAL);
                fObject = new FListView(this,lmv);
                break;
            case "HList":
                LinearLayoutManager lmh = new LinearLayoutManager(this);
                lmh.setOrientation(LinearLayoutManager.HORIZONTAL);
                fObject = new FListView(this,lmh);
                break;
            case "Dialog":
                fObject = new FDialog(this);
                break;
        }
        fObject.vid = vid;
        fObject.vtype = vtype;
        viewmap.put(vid, fObject);
        return vid;
    }

    @Override
    public String setAttr(String vid, String attr, String v1, String v2) {
        if (vid.equals("Activity")) {
            return setActivityAttr(attr,v1,v2);
        }
        FObject fObject = viewmap.get(vid);
        if (fObject != null) {
            return fObject.setAttr(attr,v1,v2);
        }
        return "";
    }

    @Override
    public String show(String vid) {
        FObject v = viewmap.get(vid);
        if (v != null) {
            rootCtn.addView(v.view);
        }
        return "";
    }

    public String newActivity(String appconf) {
        try {
            JSONObject object = (JSONObject) (new JSONTokener(appconf).nextValue());
            String launchMode = object.getString("FLaunchMode");
            String activityId = object.getString("FActivityId");
            if (launchMode == null || activityId == null) {
                return "";
            }
            Intent intent = new Intent();
            intent.putExtra("FActivityId", activityId);
            if (launchMode.equals("SingleInstance")) {
                intent.setClass(this, SingleInstanceActivity.class);
            } else if (launchMode.equals("SingleTask")) {
                intent.setClass(this, SingleTaskActivity.class);
            } else if (launchMode.equals("SingleTop")) {
                intent.setClass(this, SingleTopActivity.class);
            } else {
                intent.setClass(this, StandardActivity.class);
            }
            JSONObject intentObj = object.getJSONObject("FIntent");
            String gAction=intentObj.getString("Action");
            if (gAction!=null&&!gAction.equals("")) {
                intent.setAction(intentObj.getString("Action"));
            }
            JSONObject extraMap = intentObj.getJSONObject("Extras");
            if (extraMap != null) {
                for (Iterator<String> it = extraMap.keys(); it.hasNext(); ) {
                    String key = it.next();
                    intent.putExtra(key, extraMap.getString(key));
                }
            }
            startActivity(intent);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return "";
    }

    private String getActivityAttr(String attr) {
        switch (attr) {
            case "PackageName":
                return getPackageName();
            default:
                return "";
        }
    }

    private String setActivityAttr(String attr, final String v1, final String v2) {
        switch (attr) {
            case "RunOnUIThread":
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        Fox.triggerFunction(FoxActivity.this, v1, v2,"","");
                    }
                });
                break;
            default:
                return "";
        }
        return "";
    }

}
