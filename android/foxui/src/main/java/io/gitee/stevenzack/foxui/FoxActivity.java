package io.gitee.stevenzack.foxui;

import android.Manifest;
import android.content.BroadcastReceiver;
import android.content.ClipData;
import android.content.Context;
import android.content.Intent;
import android.content.IntentFilter;
import android.content.pm.PackageManager;
import android.graphics.drawable.Drawable;
import android.net.Uri;
import android.os.Build;
import android.os.Environment;
import android.preference.PreferenceManager;
import android.provider.Settings;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.support.v4.app.ActivityCompat;
import android.support.v4.content.LocalBroadcastManager;
import android.support.v4.graphics.drawable.DrawableCompat;
import android.support.v4.widget.DrawerLayout;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.telephony.TelephonyManager;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.Toast;

import org.json.JSONArray;
import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.AbstractSequentialList;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.concurrent.BlockingQueue;

import fox.Fox;
import fox.IActivity;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FObject.Service.FClipboard;
import io.gitee.stevenzack.foxui.FObject.Service.FService;
import io.gitee.stevenzack.foxui.FObject.Widget.FBottomNav;
import io.gitee.stevenzack.foxui.FObject.Widget.FBox;
import io.gitee.stevenzack.foxui.FObject.Widget.FButton;
import io.gitee.stevenzack.foxui.FObject.Widget.FCheck;
import io.gitee.stevenzack.foxui.FObject.Widget.FConstraintLayout;
import io.gitee.stevenzack.foxui.FObject.Widget.FCoordinatorLayout;
import io.gitee.stevenzack.foxui.FObject.Widget.FDialog;
import io.gitee.stevenzack.foxui.FObject.Widget.FEdit;
import io.gitee.stevenzack.foxui.FObject.Widget.FFab;
import io.gitee.stevenzack.foxui.FObject.Widget.FFrameLayout;
import io.gitee.stevenzack.foxui.FObject.Widget.FImage;
import io.gitee.stevenzack.foxui.FObject.Widget.FListView;
import io.gitee.stevenzack.foxui.FObject.Widget.FPopupMenu;
import io.gitee.stevenzack.foxui.FObject.Widget.FProgress;
import io.gitee.stevenzack.foxui.FObject.Widget.FRadio;
import io.gitee.stevenzack.foxui.FObject.Widget.FRadioGroup;
import io.gitee.stevenzack.foxui.FObject.Widget.FSnackbar;
import io.gitee.stevenzack.foxui.FObject.Widget.FSpace;
import io.gitee.stevenzack.foxui.FObject.Widget.FSpinner;
import io.gitee.stevenzack.foxui.FObject.Widget.FSwitch;
import io.gitee.stevenzack.foxui.FObject.Widget.FTabLayout;
import io.gitee.stevenzack.foxui.FObject.Widget.FText;
import io.gitee.stevenzack.foxui.FObject.Widget.FToolbar;
import io.gitee.stevenzack.foxui.FObject.Widget.FValueAnimator;
import io.gitee.stevenzack.foxui.FObject.Widget.FViewPager;
import io.gitee.stevenzack.foxui.FObject.Widget.FWebView;

import static io.gitee.stevenzack.foxui.Toolkit.parseMenu;

public class FoxActivity extends AppCompatActivity implements IActivity {
    private static final int FILE_SELECT_CODE = 26143;
    public Map<String, Drawable> drawableMap=new HashMap<>();
    public Map<MenuItem, String> menuItemsOnClickMap = new HashMap<>();
    public String optionMenus;
    public List<Runnable> onDestroyEvent=new ArrayList<>();
    private String activityId="MainActivity";
    public DrawerLayout rootCtn;
    public Map<String, FObject> viewmap = new HashMap<>();
    private String TAG="FoxActivity";
    public String onPermissionResults = null;
    public static final int PERMISSION_REQUEST_CODE=12351;

    public boolean notFinishFlag;
    public Map<Integer, OnActivityResultListener> onActivityResults = new HashMap<>();
    private String onBackgPressedFn;

    public interface OnActivityResultListener{
        void onActivityResult(Intent intent);
    }
    protected void mainFoxUI(DrawerLayout rootCtn) {//only called by MainActivity (the entry port of Go Code
        this.rootCtn=rootCtn;
        registerUIBro();
        Fox.main(this);
    }
    protected void startFoxUI(DrawerLayout rootCtn) {// called by those second activity
        this.rootCtn=rootCtn;
        registerUIBro();
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
    protected void onDestroy() {
        super.onDestroy();
        for (int i = 0; i < onDestroyEvent.size(); i++) {
            onDestroyEvent.get(i).run();
        }
    }

    @Override
    public void onBackPressed() {
        if (onBackgPressedFn == null) {
            super.onBackPressed();
        }
        String result=Fox.triggerFunction(this,onBackgPressedFn,"","","");
        Log.d(TAG, "onBackPressed: result="+result);
        if (result.equals("true")){
            super.onBackPressed();
        }
    }

    private void registerUIBro(){
        LocalBroadcastManager.getInstance(this).registerReceiver(new BroadcastReceiver() {
            @Override
            public void onReceive(Context context, Intent intent) {
                if (intent.getStringExtra("action").equals("quit")){
                    if (notFinishFlag) {
                        notFinishFlag = false;
                    }else {
                        finish();
                    }
                }
            }
        },new IntentFilter("uibro"));
    }
    @Override
    public String getActivityId() {
        return activityId;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        if (menuItemsOnClickMap.containsKey(item)) {
            Fox.triggerFunction(this,menuItemsOnClickMap.get(item), "","","");
            return true;
        }
        return false;
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        if (optionMenus== null || optionMenus == "" || optionMenus == "[]") {
            return false;
        }
        try {
            JSONArray array = (JSONArray) (new JSONTokener(optionMenus).nextValue());
            parseMenu(this,menu,array);
            return true;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return false;
    }

    @Override
    public String getAttr(String vid, String attr) {
        if (vid.equals("Activity")) {
            return getActivityAttr(attr);
        } else if (vid.equals("Service")) {
            return FService.getAttrStaticly(attr);
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
            case "Clipboard":
                fObject = new FClipboard(this);
                break;
            case "Edit":
                fObject = new FEdit(this);
                break;
            case "PopupMenu":
                try {
                    String[] vids = vid.split(":");
                    FPopupMenu popupMenu = new FPopupMenu(this, viewmap.get(vids[1]));
                    popupMenu.vtype = vtype;
                    popupMenu.vid = vids[0];
                    viewmap.put(popupMenu.vid, popupMenu);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                return "";
            case "Progress":
                fObject = new FProgress(this);
                break;
            case "Radio":
                fObject = new FRadio(this);
                break;
            case "RadioGroup":
                fObject = new FRadioGroup(this);
                break;
            case "ViewPager":
                fObject = new FViewPager(this);
                break;
            case "TabLayout":
                fObject = new FTabLayout(this);
                break;
            case "Fab":
                fObject = new FFab(this);
                break;
            case "Snackbar":
                String[] vids=vid.split(",");
                fObject = new FSnackbar(this, viewmap.get(vids[1]).view);
                fObject.vid = vids[0];
                fObject.vtype = vtype;
                viewmap.put(fObject.vid, fObject);
                return fObject.vid;
            case "CoordinatorLayout":
                fObject = new FCoordinatorLayout(this);
                break;
            case "Space":
                fObject = new FSpace(this);
                break;
            case "Spinner":
                fObject = new FSpinner(this);
                break;
            case "Switch":
                fObject = new FSwitch(this);
                break;
            case "BottomNav":
                fObject = new FBottomNav(this);
                break;
            case "Toolbar":
                fObject = new FToolbar(this);
                break;
            case "ValueAnimator":
                fObject = new FValueAnimator(this);
                break;
            case "WebView":
                fObject = new FWebView(this);
                break;
            case "Check":
                fObject = new FCheck(this);
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
        } else if (vid.equals("Service")) {
            FService.setAttr(this,attr, v1);
            return "";
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
        if (v.vtype.equals("Snackbar")) {
            v.setAttr("Show", "", "");
        }else  {
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
            case "IMEI":
                TelephonyManager telephonyManager = (TelephonyManager) getSystemService(Context.TELEPHONY_SERVICE);
                if (ActivityCompat.checkSelfPermission(this, Manifest.permission.READ_PHONE_STATE) != PackageManager.PERMISSION_GRANTED) {
                    return Settings.Secure.getString(getContentResolver(), Settings.Secure.ANDROID_ID);
                }
                return telephonyManager.getDeviceId();
            case "UniqueID":
                String android_id = Settings.Secure.getString(getContentResolver(),
                        Settings.Secure.ANDROID_ID);
                return android_id;
            case "Build.MODEL":
                return Build.MODEL;
            case "ExternalStorageDirectory":
                return Environment.getExternalStorageDirectory().getAbsolutePath();
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
            case "CheckSelfPermission":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    return String.valueOf(checkSelfPermission(v1)==PackageManager.PERMISSION_GRANTED);
                }
                break;
            case "RequestPermissions":
                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
                    String[] strs=v1.split(":");
                    onPermissionResults=v2;
                    requestPermissions(strs,PERMISSION_REQUEST_CODE);
                }
                break;
            case "ShowToast":
                Toast.makeText(this,v1,Toast.LENGTH_SHORT).show();
                break;
            case "BackPressed":
                onBackPressed();
                break;
            case "OnBackPressed":
                onBackgPressedFn = v1;
                break;
            case "Finish":
                finish();
                break;
            case "OpenFileChooser":// [ type : "*/*" , allowMultiple : "true" , callback : "218hxjgs861h9cb1298" ]
                try {
                    Intent intent1 = new Intent(Intent.ACTION_GET_CONTENT);
                    final JSONArray array = (JSONArray) (new JSONTokener(v1).nextValue());
                    intent1.setType(array.getString(0));
                    intent1.addCategory(Intent.CATEGORY_OPENABLE);
                    intent1.putExtra(Intent.EXTRA_ALLOW_MULTIPLE, array.getString(1).equals("true"));
                    startActivityForResult(Intent.createChooser(intent1, ""),FILE_SELECT_CODE);
                    final String fnID = array.getString(2);
                    onActivityResults.put(FILE_SELECT_CODE, new OnActivityResultListener() {
                        @Override
                        public void onActivityResult(Intent data) {
                            if (data == null) {
                                return;
                            }
                            Uri uri = data.getData();
                            JSONArray array1 = new JSONArray();
                            try {
                                if (uri != null) {
                                    String path = null;
                                    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT) {
                                        path = Toolkit.getPathByUri4kitkat(FoxActivity.this, uri);
                                    }else {
                                        path=Toolkit.getPathByUriBelowKitkat(FoxActivity.this,uri);
                                    }
                                    if (path != null) {
                                        array1.put(path);
                                    }
                                } else {
                                    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
                                        ClipData clipData = data.getClipData();
                                        if (clipData != null) {
                                            for (int i=0;i<clipData.getItemCount();i++) {
                                                ClipData.Item item = clipData.getItemAt(i);
                                                Uri uri1=item.getUri();
                                                String url = null;
                                                if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT) {
                                                    url = Toolkit.getPathByUri4kitkat(FoxActivity.this, uri1);
                                                }else {
                                                    url=Toolkit.getPathByUriBelowKitkat(FoxActivity.this,uri1);
                                                }
                                                if (url!=null)
                                                    array1.put(url);
                                            }
                                        }
                                    }else {
                                        Toast.makeText(FoxActivity.this,"No file manager installed",Toast.LENGTH_LONG).show();
                                    }
                                }
                                if (array1.length() == 0) {
                                    Toast.makeText(FoxActivity.this,"Failed",Toast.LENGTH_LONG).show();
                                }else {
                                    Fox.triggerFunction(FoxActivity.this,fnID, array1.toString(),"","");
                                }
                            }catch (Exception e){
                                e.printStackTrace();
                                Toast.makeText(FoxActivity.this,e.toString(),Toast.LENGTH_LONG).show();
                            }
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                    Toast.makeText(FoxActivity.this,e.toString(),Toast.LENGTH_LONG).show();
                }
                break;
            default:
                return "";
        }
        return "";
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, @Nullable Intent data) {
        if (onActivityResults.containsKey(requestCode)) {
            onActivityResults.get(requestCode).onActivityResult(data);
        }
    }
    @Override
    public void onRequestPermissionsResult(int requestCode, @NonNull String[] permissions, @NonNull int[] grantResults) {
        if (onPermissionResults != null) {
            JSONArray bs = new JSONArray();
            for (int i = 0; i < grantResults.length; i++) {
                bs.put(grantResults[i] == PackageManager.PERMISSION_GRANTED);
            }
            Fox.triggerFunction(this,onPermissionResults,bs.toString(),"","");
            onPermissionResults = null;
        }
    }
}
