package com.example.demo.controller;

import com.example.demo.api.CommonResult;
import com.example.demo.service.impl.SparrowGodServiceImpl;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.ApiParam;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Api("雀神接口")
@RestController
@RequestMapping("/sparrow")
public class SparrowGodController {

    private final SparrowGodServiceImpl sparrowGodService;

    @Autowired
    public SparrowGodController(SparrowGodServiceImpl sparrowGodService) {
        this.sparrowGodService = sparrowGodService;
    }


    @ApiOperation("查看空余房间")
    @RequestMapping(value = "/claim", method = RequestMethod.GET)
    public CommonResult<List<Integer>> claim() {
        List<Integer> rooms = sparrowGodService.claimRoom();

        return CommonResult.success(rooms);
    }

    @ApiOperation("购买房间")
    @RequestMapping(value = "/join/{id}", method = RequestMethod.GET)
    public CommonResult<String> join(@PathVariable(name = "id") @ApiParam(value = "房间号") int id) {
        return CommonResult.success(String.valueOf(id));
    }

    @ApiOperation("加入成员")
    @RequestMapping("/sit/{id}")
    public CommonResult<String> sit(
            @PathVariable(name = "id") @ApiParam(value = "房间号") int id,
            @RequestParam(name = "member") List<String> members) {
        return CommonResult.success("ss");
    }

    @ApiOperation("对局开始")
    @RequestMapping("/start/{id}")
    public CommonResult<String> start(
            @PathVariable(name = "id") @ApiParam(value = "房间号") int id,
            @RequestParam(name = "member") List<String> members) {
        return CommonResult.success("ss");
    }

    @ApiOperation("录入胜利者")
    @RequestMapping("/win/{id}/{member}")
    public CommonResult<String> win(
            @PathVariable(name = "id") @ApiParam(value = "房间号") int id,
            @PathVariable(name = "member") String member
                                      ) {
        return CommonResult.success("ss");
    }

    @ApiOperation("结算")
    @RequestMapping("/settle_accounts/{id}")
    public CommonResult<String> settle_accounts(
            @PathVariable(name = "id") @ApiParam(value = "房间号") int id
    ) {
        return CommonResult.success("ss");
    }

    @ApiOperation("对局结束")
    @RequestMapping("/end/{id}")
    public CommonResult<String> end(
            @PathVariable(name = "id") @ApiParam(value = "房间号") int id) {
        return CommonResult.success("ss");
    }

    @ApiOperation("退出房间")
    @RequestMapping("/exit/{id}")
    public CommonResult<String> exit(
            @PathVariable(name = "id") @ApiParam(value = "房间号") int id) {
        return CommonResult.success("ss");
    }
}
