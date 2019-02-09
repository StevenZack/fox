package io.gitee.stevenzack.foxui.FObject.Widget;

import android.graphics.Color;
import android.graphics.drawable.Drawable;
import android.support.v4.view.ViewCompat;
import android.support.v7.widget.Toolbar;
import android.view.View;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;
import io.gitee.stevenzack.foxui.FoxActivity;
import io.gitee.stevenzack.foxui.R;
import io.gitee.stevenzack.foxui.Toolkit;

public class FToolbar extends FObject {
    public Toolbar v;
    private String navigationIconSrc;

    public FToolbar(FoxActivity foxActivity) {
        parentController = foxActivity;
        v = (Toolbar) parentController.getLayoutInflater().inflate(R.layout.my_toolbar, parentController.rootCtn, false);
        v.setId(ViewCompat.generateViewId());
        view=v;
        setElevation("4");
        parentController.setSupportActionBar(v);
        parseSize("-2","-1");
    }

    @Override
    public String getAttr(String attr) {
        String str = super.getAttr(attr);
        if (str != null) {
            return str;
        }
        switch (attr){
            //----------------------------------------------
            case "Title":
                return (String) v.getTitle();
            case "SubTitle":
                return (String) v.getSubtitle();
        }
        return "";
    }

    @Override
    public String setAttr(String attr, final String value, String value2) {
        String str = super.setAttr(attr, value, value2);
        if (str != null) {
            return str;
        }
        switch (attr) {
            case "Title":
                parentController.getSupportActionBar().setTitle(value);
                break;
            case "TitleColor":
                v.setTitleTextColor(Color.parseColor(value));
                break;
            case "SubTitle":
                parentController.getSupportActionBar().setSubtitle(value);
                break;
            case "SubTitleColor":
                v.setSubtitleTextColor(Color.parseColor(value));
                break;
            case "Menus":
                parentController.optionMenus=value;
                break;
            case "BackNavigation":
                if (value.equals("true")) {
                    v.setNavigationOnClickListener(new View.OnClickListener() {
                        @Override
                        public void onClick(View v) {
                            parentController.onBackPressed();
                        }
                    });
                }
                parentController.getSupportActionBar().setDisplayHomeAsUpEnabled(value.equals("true"));
                break;
            case "NavigationIcon":
                if (navigationIconSrc != null && value.equals(navigationIconSrc)) {
                    break;
                }
                Toolkit.file2Drawable(parentController, value, new Toolkit.OnDrawableReadyListener() {
                    @Override
                    public void onDrawableReady(Drawable drawable) {
                        v.setNavigationIcon(drawable);
                    }
                });
                break;
            case "OnNavigationIconClick":
                v.setNavigationOnClickListener(new View.OnClickListener() {
                    @Override
                    public void onClick(View v) {
                        Fox.triggerFunction(parentController, value, "", "", "");
                    }
                });
                break;
        }
        return "";
    }
}
