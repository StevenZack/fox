package io.gitee.stevenzack.foxui.FObject.Widget;


import android.os.Bundle;
import android.support.v4.app.Fragment;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import fox.Fox;
import io.gitee.stevenzack.foxui.FObject.FObject;

public  class DemoObjectFragment extends Fragment {
    public static final String ARG_OBJECT = "object";
    private FViewPager fviewPager;

    public DemoObjectFragment setFviewPager(FViewPager fviewPager) {
        this.fviewPager=fviewPager;
        return this;
    }
    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
        Bundle args = getArguments();
        int i=args.getInt(ARG_OBJECT);
        if (fviewPager == null) {
            return null;
        }
        if (fviewPager.onCreateView != null && !fviewPager.onCreateView.equals("")) {
            Log.d("TAG", "onCreateView: 1");
            String vid = Fox.triggerFunction(fviewPager.parentController,fviewPager.onCreateView, String.valueOf(i),"","");
            if (fviewPager.parentController.viewmap.containsKey(vid)) {
                Log.d("TAG", "onCreateView: 2");
                FObject cview = fviewPager.parentController.viewmap.get(vid);
                if (cview != null) {
                    Log.d("TAG", "onCreateView: 3");
                    return cview.view;
                }
            }
        }
        if (fviewPager.pages.size() <= i) {
            return null;
        }
        FViewPager.FPage p = fviewPager.pages.get(i);
        View rootView=fviewPager.parentController.viewmap.get(Fox.triggerFunction(fviewPager.parentController,p.vid,"","","")).view;
        return rootView;
    }
}
