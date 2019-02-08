package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.support.design.widget.TabLayout;
import android.support.v4.view.ViewCompat;

import org.json.JSONObject;
import org.json.JSONTokener;

import java.util.ArrayList;
import java.util.List;

import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.R;
import io.gitee.stevenzack.foxui.Toolkit;

public class FTabLayout extends FObject {
    public TabLayout v;
    public List<String> tabsList = new ArrayList<>();

    public FTabLayout(FoxActivity activity) {
        parentController = activity;
        v = new TabLayout(activity);
        v.setId(ViewCompat.generateViewId());
        view=v;
        setElevation("4");
        parseSize("-2","-1");
        v.setTabTextColors(Color.parseColor("#dddddd"),Color.WHITE);
        v.setSelectedTabIndicatorColor(parentController.getResources().getColor(R.color.colorAccent));
        v.setBackgroundColor(parentController.getResources().getColor(R.color.colorPrimary));
    }
    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            //----------------------------------------------
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
            case "TabTextColors":
                    v.setTabTextColors(Color.parseColor(value),Color.parseColor(value2));
                break;
            case "TabIndicatorColor":
                try {
                    v.setSelectedTabIndicatorColor(Color.parseColor(value));
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "AddTab":
                try {
                    JSONObject object = (JSONObject) (new JSONTokener(value).nextValue());
                    String icon = object.getString("Icon");
                    String text = object.getString("Text");
                    final TabLayout.Tab tab=v.newTab();
                    tabsList.add(text);
                    tab.setText(text);
                    if (!icon.equals("")) {
                        Toolkit.file2Drawable(parentController, icon, new Toolkit.OnDrawableReadyListener() {
                            @Override
                            public void onDrawableReady(Drawable drawable) {
                                tab.setIcon(drawable);
                            }
                        });
                    }
                    v.addTab(tab);
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "ViewPager":
                Object o=parentController.viewmap.get(value);
                if (o==null) break;
                FViewPager fViewPager = (FViewPager) o;
                v.setupWithViewPager(fViewPager.v);
                break;
        }
        return "";
    }
}
