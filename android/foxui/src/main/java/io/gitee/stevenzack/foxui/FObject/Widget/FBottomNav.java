package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.drawable.Drawable;
import android.os.Build;
import android.support.annotation.NonNull;
import android.support.design.widget.BottomNavigationView;
import android.support.v4.view.ViewCompat;
import android.util.TypedValue;
import android.view.MenuItem;

import org.json.JSONArray;
import org.json.JSONTokener;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.Toolkit;

public class FBottomNav extends FObject {
    public BottomNavigationView v;

    public FBottomNav(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = new BottomNavigationView(parentController);
        v.setId(ViewCompat.generateViewId());
        view=v;
        parseSize("-2","-1");

        TypedValue outValue = new TypedValue();
        parentController.getTheme().resolveAttribute(android.R.attr.windowBackground, outValue, true);
        Drawable d= parentController.getResources().getDrawable(outValue.resourceId);
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN) {
            v.setBackground(d);
        }else {
            v.setBackgroundDrawable(d);
        }
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
        if (str!=null){
            return str;
        }
        switch (attr) {
            case "Menus":
                try {
                    JSONArray object = (JSONArray) (new JSONTokener(value).nextValue());
                    Toolkit.parseMenu(parentController, v.getMenu(), object);
                    v.setOnNavigationItemSelectedListener(new BottomNavigationView.OnNavigationItemSelectedListener() {
                        @Override
                        public boolean onNavigationItemSelected(@NonNull MenuItem menuItem) {
                            if (parentController.menuItemsOnClickMap.containsKey(menuItem)) {
//                                Faithdroid.triggerEventHandler(parentController.menuItemsOnClickMap.get(menuItem), "");
                                Fox.triggerFunction(parentController, parentController.menuItemsOnClickMap.get(menuItem), "", "", "");
                                return true;
                            }
                            return false;
                        }
                    });
                } catch (Exception e) {
                    e.printStackTrace();
                }
                break;
            case "SelectedIndex":
                v.setSelectedItemId(v.getMenu().getItem(Integer.parseInt(value)).getItemId());
                break;
        }
        return "";
    }
}
